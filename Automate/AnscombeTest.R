#!/Program Files/R/R-4.3.1/bin/x64/Rgui Rscript


args = commandArgs(trailingOnly=TRUE)
set <- args[1]
run_n <- 500

if(!require(jsonlite)){
  install.packages("jsonlite", repos = 'https://cran.r-project.org')
  library(jsonlite)
}

#load data
anscombe <- list(
  One = data.frame(
		x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
		y = c(8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68)
    ),
  Two = data.frame(
		x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
		y = c(9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74)
    ),
	Three = data.frame(
		x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
		y = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73)
    ),
	Four = data.frame(
		x = c(8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8),
		y = c(6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89)
	  )
  )

experiment <- function(dataset){
    model = lm(y ~ x, data = anscombe[[set]])
    intercept <- summary(model)$coefficients["(Intercept)", "Estimate"]
    slope <- summary(model)$coefficients["x", "Estimate"]
    values = c(intercept, slope)
}

n = run_n
start_time <- Sys.time()
for (i in 1:n) {
    experiment(dataset = set)
}
end_time <- Sys.time()
runtime <- end_time - start_time
averageTime = as.numeric(runtime/n)
jsonlite::toJSON(list(Coefficients = experiment(dataset = set), Time = averageTime), pretty = TRUE, digits = 20, auto_unbox = TRUE)