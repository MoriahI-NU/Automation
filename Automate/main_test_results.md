$ go test main_test.go main.go  
  

--- FAIL: TestExperimentPython (1.43s)  

    --- FAIL: TestExperimentPython/Test_Case_Python (1.43s)  

        main_test.go:27: ExperimentPython() = {{[3.0000909090909085, 0.5000909090909091] 0.0001875}  

            }, want {[3, 0.5] 0.0002}  
  

--- FAIL: TestExperimentR (0.53s)  

    --- FAIL: TestExperimentR/Test_Case_R (0.53s)  

        main_test.go:51: ExperimentR() = {{  

              [3.0000909090909089, 0.50009090909090914],  

               0.00069842576980590817  

            }  

            }, want {[3, 0.5] 0.0006}  

--- FAIL: TestExperimentGo (0.01s)  
  
  
    --- FAIL: TestExperimentGo/Test_Case_Go (0.01s)  

        main_test.go:78: ExperimentGo() = {3.0000909090909014 0.5000909090909129 0.000018300999999999996}, want {[3, 0.5] 0.0001}  


FAIL  

FAIL    command-line-arguments  2.068s  

FAIL