# Trading Arbitrage Detector
This project implements a trading arbitrage detector using graph theory and the Bellman-Ford algorithm. Arbitrage in financial markets involves taking advantage of price differences of the same asset across different markets to make a profit. This tool helps identify such arbitrage opportunities by modeling exchange rates as a graph and detecting negative weight cycles.

## Bellman-Ford Algorithm

The Bellman-Ford algorithm is used to find the shortest paths from a single source vertex to all other vertices in a weighted graph. It can handle graphs with negative weights and detect negative weight cycles. If a shorter path is found after V-1 iterations (where V is the number of vertices), it indicates a negative weight cycle.

## Project Structure

- `main.go`: Example usage of the Graph structure and detecting arbitrage.
- `graph.go`: Contains the Graph structure, Edge structure, and methods for adding edges and detecting arbitrage.
- `graph_test.go`: Contains unit tests for the Graph methods.

### Getting Started

1. **Clone the repository**:
   ```sh
   git clone https://github.com/mdshahjahanmiah/trading-arbitrage-detector.git
   cd trading-arbitrage-detector
   ```

2. **Run and Test the application:**:
   ```sh
    go mod tidy
    go run cmd/main.go
    go test ./graph
   ```



