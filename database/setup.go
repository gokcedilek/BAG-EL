package database

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"project/util"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func createTable(
	context context.Context, svc *dynamodb.Client,
	tableDefinition *dynamodb.CreateTableInput,
) {
	out, err := svc.CreateTable(context, tableDefinition)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created database", out)
}

func CreateTableIfNotExists(svc *dynamodb.Client, tableName string) {
	tables, _ := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{})

	tableExists := false

	for _, name := range tables.TableNames {
		if name == tableName {
			tableExists = true
			break
		}
	}

	if tableExists {
		return
	}

	CreateTable(context.TODO(), svc, tableName)
}

func CreateTable(
	context context.Context, svc *dynamodb.Client, tableName string,
) {
	bagelDefinition := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	}

	createTable(context, svc, bagelDefinition)
	waitForTable(context, svc, tableName)
}

func AddGraph(svc *dynamodb.Client, filePath string, tableName string) {
	graph := ParseInputGraph(filePath, true)
	vertices := graphToVertices(graph)
	batches := CreateBatches(vertices)
	BatchInsertVertices(svc, tableName, batches)
}

func waitForTable(ctx context.Context, db *dynamodb.Client, tn string) error {
	w := dynamodb.NewTableExistsWaiter(db)
	err := w.Wait(
		ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tn),
		},
		2*time.Minute,
		func(o *dynamodb.TableExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		},
	)
	if err != nil {
		panic(err)
	}

	return err
}

func ParseInputGraph(filePath string, isCSV bool) map[uint64][]uint64 {
	graph := make(map[uint64][]uint64)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	delimiter := "\t"
	if isCSV {
		delimiter = ","
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "#") {
			continue
		}

		edge := strings.Split(line, delimiter)
		src, _ := strconv.ParseUint(edge[0], 10, 32)
		dest, _ := strconv.ParseUint(edge[1], 10, 32)

		graph[uint64(src)] = append(graph[uint64(src)], uint64(dest))
		if graph[uint64(dest)] == nil {
			graph[uint64(dest)] = []uint64{}
		}
	}
	fmt.Printf("Successfully parsed %v nodes\n", len(graph))
	return graph
}

func WriteGraphToFile(fileName string, graph map[uint64][]uint64) error {
	f, err := os.Create(fileName)
	writer := bufio.NewWriter(f)

	if err != nil {
		return err
	}

	for src, edges := range graph {
		for _, e := range edges {
			_, err := writer.WriteString(fmt.Sprintf("%d,%d\n", src, e))
			if err != nil {
				return err
			}
		}
		writer.Flush()
	}

	return nil
}

func ReduceGraphToXNodes(graph map[uint64][]uint64, desiredNumNodes int, degree int) map[uint64][]uint64 {
	if len(graph) < desiredNumNodes {
		fmt.Errorf("graph has %d nodes which is less than the desired number of %d nodes",
			len(graph), desiredNumNodes)
		return nil
	}

	reducedGraph := make(map[uint64][]uint64)
	deletedNodes := make(map[uint64]bool)
	n := len(graph)
	numNodesToBeDeleted := n - desiredNumNodes

	// delete first n nodes
	for v, edges := range graph {
		if isDeletedNode(v, deletedNodes) {
			continue
		}
		deletedNodes[v] = true
		numNodesToBeDeleted--

		for _, e := range edges {
			if !isDeletedNode(e, deletedNodes) {
				deletedNodes[e] = true
				numNodesToBeDeleted--
				if numNodesToBeDeleted == 0 {
					break
				}
			}
		}

		if numNodesToBeDeleted == 0 {
			break
		}
	}
	// copy nodes that aren't deleted & delete edges that exist
	for v, edges := range graph {
		if isDeletedNode(v, deletedNodes) {
			continue
		}
		n := len(edges)
		reducedGraph[v] = make([]uint64, n, n)
		copy(reducedGraph[v], edges)
		reducedGraph[v] = removeEdgesWithDeletedDest(reducedGraph[v], deletedNodes)
	}

	PostProcessReducedGraph(reducedGraph, degree)

	return reducedGraph
}

func PostProcessReducedGraph(graph map[uint64][]uint64, degree int) {
	edges := getSliceOfEdges(graph)
	n := len(edges)

	for edge := range graph {
		currDegree := len(graph[edge])
		if degree <= currDegree {
			continue
		}

		requiredDegree := degree - currDegree
		newDegree := rand.Int() % requiredDegree
		if newDegree == 0 {
			newDegree += 1
		}
		mapOfEdges := getMapOfOutgoingEdges(graph[edge])
		for i := 0; i < newDegree; i++ {
			newEdge := edges[rand.Int()%n]
			if _, ok := mapOfEdges[newEdge]; !ok {
				graph[edge] = append(graph[edge], newEdge)
				mapOfEdges[newEdge] = true
			} else {
				i--
			}
		}
	}
}

func getMapOfOutgoingEdges(edges []uint64) map[uint64]bool {
	mapOfEdges := make(map[uint64]bool)
	for _, edge := range edges {
		mapOfEdges[edge] = true
	}
	return mapOfEdges
}

func getSliceOfEdges(graph map[uint64][]uint64) []uint64 {
	edges := make([]uint64, len(graph))
	i := 0
	for edge := range graph {
		edges[i] = edge
	}

	return edges
}

func removeEdge(edges []uint64, idx int) []uint64 {
	n := len(edges)
	edges[idx] = edges[n-1]
	return edges[:n-1]
}

func removeEdgesWithDeletedDest(edges []uint64, deletedNodes map[uint64]bool) []uint64 {
	numIterations, idx := len(edges), 0

	for i := 0; i < numIterations; i++ {
		if isDeletedNode(edges[idx], deletedNodes) {
			edges = removeEdge(edges, idx)
			idx--
		}
		idx++
	}

	return edges
}

func isDeletedNode(node uint64, deleted map[uint64]bool) bool {
	_, ok := deleted[node]
	return ok
}

func graphToVertices(graph Graph) []Vertex {
	vertices := make([]Vertex, len(graph))

	idx := 0
	for vertexId, edges := range graph {
		hash := util.HashId(vertexId)
		vertices[idx] = Vertex{
			ID:    vertexId,
			Edges: edges,
			Hash:  hash,
		}
		idx++
	}

	return vertices
}
