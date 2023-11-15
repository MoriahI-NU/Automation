$ go test main_test.go main.go  

--- FAIL: TestExperimentPython (1.41s)  

    --- FAIL: TestExperimentPython/Test_Case_Python (1.41s)  

        main_test.go:25: ExperimentPython() = {{"Coefficients": [3.0000909090909085, 0.5000909090909091], "Time": 0.00021875}
             <nil>}, want {{"Coefficients": [3, 0.5], "Time": 0.0002 <nil>}  


--- FAIL: TestExperimentR (0.53s)  

    --- FAIL: TestExperimentR/Test_Case_R (0.53s)  

        main_test.go:46: ExperimentR() = {{  
              "Coefficients": [3.0000909090909089, 0.50009090909090914],  
              "Time": 0.00069193220138549804  
            }  
             <nil>}, want {{"Coefficients": [3, 0.5], "Time": 0.0006} <nil>}  
0.0000190  


--- FAIL: TestExperimentGo (0.01s)  

    --- FAIL: TestExperimentGo/Test_Case_Go (0.01s)  

        main_test.go:67: ExperimentGo() = {[3.0000909090909014 0.5000909090909129] 1.8973600000000004e-05}, want {[3 0.5] 9.1e-07}  

FAIL  

FAIL    command-line-arguments  2.049s  

FAIL