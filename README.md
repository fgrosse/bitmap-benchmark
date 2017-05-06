```bash
BenchmarkIntersectionRoaring-4           	 1000000	      1295 ns/op
BenchmarkIntersectionPilosa-4            	    2000	   1124238 ns/op
BenchmarkIntersectionPilosaInternal-4    	    1000	   1470149 ns/op

BenchmarkUnionRoaring-4                  	    5000	    320838 ns/op
BenchmarkUnionPilosa-4                   	    1000	   1791906 ns/op
BenchmarkUnionPilosaInternal-4           	    1000	   2157900 ns/op

BenchmarkSetRoaring-4                    	20000000	        70.2 ns/op
BenchmarkSetPilosa-4                     	20000000	        92.6 ns/op
BenchmarkSetPilosaInternal-4             	20000000	        71.6 ns/op

BenchmarkGetTestRoaring-4                	10000000	       141 ns/op
BenchmarkGetTestPilosa-4                 	     500	   3201099 ns/op
BenchmarkGetTestPilosaInternal-4         	50000000	        38.2 ns/op

BenchmarkCountRoaring-4                  	30000000	        51.0 ns/op
BenchmarkCountPilosa-4                   	1000000000	         2.60 ns/op
BenchmarkCountPilosaInternal-4           	100000000	        12.4 ns/op

BenchmarkIterateRoaring-4                	    2000	    870095 ns/op
BenchmarkIteratePilosaInternal-4         	    2000	   1033366 ns/op

BenchmarkSparseIterateRoaring-4          	    2000	    878023 ns/op
BenchmarkSparseIteratePilosaInternal-4   	    3000	    544532 ns/op
```
