.PHONY: bench ballocopt ballocnonopt ballocsfs ballocsnfs ballocafs ballocanfs benchstat fieldalignment lint

# Запускает все бенчи последовательно и сохраняет вывод
bench:
	$(MAKE) ballocopt
	$(MAKE) ballocnonopt
	$(MAKE) ballocsfs
	$(MAKE) ballocsnfs
	$(MAKE) ballocafs
	$(MAKE) ballocanfs
	$(MAKE) benchstat

# Бенч только Alloc (Optimized)
ballocopt:
	go test -bench ^BenchmarkAllocOptimized$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocOptimized.txt

# Бенч только Alloc (NonOptimized)
ballocnonopt:
	go test -bench ^BenchmarkAllocNonOptimized$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocNonOptimized.txt

# Бенч только BenchmarkAllocSimpleFalseSharing
ballocsfs:
	go test -bench ^BenchmarkAllocSimpleFalseSharing$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocSimpleFalseSharing.txt

# Бенч только BenchmarkAllocSimpleNoFalseSharing
ballocsnfs:
	go test -bench ^BenchmarkAllocSimpleNoFalseSharing$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocSimpleNoFalseSharing.txt

# Бенч только BenchmarkAllocAtomicFalseSharing
ballocafs:
	go test -bench ^BenchmarkAllocAtomicFalseSharing$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocAtomicFalseSharing.txt

# Бенч только BenchmarkAllocAtomicNoFalseSharing
ballocanfs:
	go test -bench ^BenchmarkAllocAtomicNoFalseSharing$$ -benchmem ./... -benchtime=10s -shuffle on -count=10 > BenchmarkAllocAtomicNoFalseSharing.txt


# Сравнение результатов через benchstat (нужен go install golang.org/x/perf/cmd/benchstat)
benchstat:
	go install golang.org/x/perf/cmd/benchstat@latest
	benchstat BenchmarkAllocOptimized.txt BenchmarkAllocNonOptimized.txt > bench_stat_simple.txt
	benchstat BenchmarkAllocSimpleFalseSharing.txt BenchmarkAllocSimpleNoFalseSharing.txt > bench_stat_current_simple.txt
	benchstat BenchmarkAllocAtomicFalseSharing.txt BenchmarkAllocAtomicNoFalseSharing.txt > bench_stat_current_atomic.txt

## LINTERS
.PHONY: fieldalignment
fieldalignment:
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	fieldalignment ./...

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml