# pdpr-go: Cryptographic Algorithms Benchmark Suite

This project implements and benchmarks several cryptographic algorithms, including hash functions, encryption schemes, and homomorphic hiding mechanisms. The benchmarks are written in Go and can be executed using the Go toolchain.

## Project Structure

- `crypto/hash/ghash/` — GHash cryptographic hash function and benchmarks
- `crypto/hash/lthash/` — LtHash cryptographic hash function and benchmarks
- `crypto/encryption/gcrypt/` — GCrypt encryption scheme and benchmarks
- `crypto/homomorphic_hiding/dlhh/` — DLHH homomorphic hiding and benchmarks

## Running Benchmarks

All benchmarks are implemented as Go benchmark tests (functions starting with `Benchmark`) in files ending with `_bench_test.go`. To run the benchmarks, use the following commands from the project root:

### 1. GHash Benchmarks

```
go test -bench=. ./crypto/hash/ghash/
```
This will run all benchmarks in `ghash_bench_test.go` and `pdpr_bench_test.go`, including:
- GHash performance for various parameters
- Full PDPr protocol benchmarks (encryption, proof generation, verification, decryption)

### 2. LtHash Benchmarks

```
go test -bench=. ./crypto/hash/lthash/
```
This will run all benchmarks in `lthash_bench_test.go`, including LtHash performance for different file and block sizes.

### 3. GCrypt Encryption Benchmarks

```
go test -bench=. ./crypto/encryption/gcrypt/
```
This will run all benchmarks in `gcrypt_bench_test.go`, including encryption and decryption performance for various key and message sizes.

### 4. DLHH Homomorphic Hiding Benchmarks

```
go test -bench=. ./crypto/homomorphic_hiding/dlhh/
```
This will run all benchmarks in `dlhh_bench_test.go`, including homomorphic hiding, encryption, decryption, proof generation, and verification.

## Customizing Benchmark Runs

You can pass additional flags to control the benchmarks, for example:
- `-benchtime=5s` to increase the benchmark duration
- `-count=3` to repeat benchmarks multiple times

Example:
```
go test -bench=Benchmark__Encrypt -benchtime=5s -count=3 ./crypto/encryption/gcrypt/
```

## Requirements
- Go 1.18 or newer

## Notes
- Benchmarks may require significant memory and CPU, especially for large file or key sizes.
- Results are reported in milliseconds per operation (e.g., `ms/hash`, `ms/encryption`).

## References
- See the source code in each subdirectory for details on the algorithms and their parameters.
