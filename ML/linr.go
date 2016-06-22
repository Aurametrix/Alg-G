package main

import "fmt"

func calc_ols_params(y []float64, x[][]float64, n_iterations int, alpha float64) []float64 {

    thetas := make([]float64, len(x))

    for i := 0; i < n_iterations; i++ {

        my_diffs := calc_diff(thetas, y, x)

        my_grad := calc_gradient(my_diffs, x)

        for j := 0; j < len(my_grad); j++ {
            thetas[j] += alpha * my_grad[j]
        }
    }
    return thetas
}

func calc_diff (thetas []float64, y []float64, x[][]float64) []float64 {
    diffs := make([]float64, len(y))
    for i := 0; i < len(y); i++ {
        prediction := 0.0
        for j := 0; j < len(thetas); j++ {
            prediction += thetas[j] * x[j][i]
        }
        diffs[i] = y[i] - prediction
    }
    return diffs
}

func calc_gradient(diffs[] float64, x[][]float64) []float64 {
    gradient := make([]float64, len(x))
    for i := 0; i < len(diffs); i++ {
        for j := 0; j < len(x); j++ {
            gradient[j] += diffs[i] * x[j][i]
        }
    }
    for i := 0; i < len(x); i++ {
        gradient[i] = gradient[i] / float64(len(diffs))
    }

    return gradient
}

func main(){
    y := []float64 {3,4,5,6,7}
    x := [][]float64 {{1,1,1,1,1}, {4,3,2,1,3}}

    thetas := calc_ols_params(y, x, 100000, 0.001)

    fmt.Println("Thetas : ", thetas)

    y_2 := []float64 {1,2,3,4,3,4,5,4,5,5,4,5,4,5,4,5,6,5,4,5,4,3,4}

    x_2 := [][]float64 {{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
                            {4,2,3,4,5,4,5,6,7,4,8,9,8,8,6,6,5,5,5,5,5,5,5},
                    {4,1,2,3,4,5,6,7,5,8,7,8,7,8,7,8,7,7,7,7,7,6,5},
                    {4,1,2,5,6,7,8,9,7,8,7,8,7,7,7,7,7,7,6,6,4,4,4},}

    thetas_2 := calc_ols_params(y_2, x_2, 100000, 0.001)

    fmt.Println("Thetas_2 : ", thetas_2)

}
