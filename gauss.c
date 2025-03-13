#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void eliminacaoGauss(int n, double **matriz, double *solucao) {
    int i, j, k, linha_pivo;
    double max_val, fator;

    for (i = 0; i < n; i++) {
        linha_pivo = i;
        max_val = fabs(matriz[i][i]);
        for (j = i + 1; j < n; j++) {
            if (fabs(matriz[j][i]) > max_val) {
                max_val = fabs(matriz[j][i]);
                linha_pivo = j;
            }
        }

        if (linha_pivo != i) {
            for (k = 0; k <= n; k++) {
                double temp = matriz[i][k];
                matriz[i][k] = matriz[linha_pivo][k];
                matriz[linha_pivo][k] = temp;
            }
        }

        if (fabs(matriz[i][i]) < 1e-10) {
            printf("Erro: Matriz singular, sem solução única.\n");
            exit(EXIT_FAILURE);
        }

        for (j = i + 1; j < n; j++) {
            fator = matriz[j][i] / matriz[i][i];
            for (k = i; k <= n; k++) {
                matriz[j][k] -= fator * matriz[i][k];
            }
        }
    }

    for (i = n - 1; i >= 0; i--) {
        solucao[i] = matriz[i][n];
        for (j = i + 1; j < n; j++) {
            solucao[i] -= matriz[i][j] * solucao[j];
        }
        solucao[i] /= matriz[i][i];
    }
}

int main() {
    int n, i, j;
    
    printf("Numero de variaveis: ");
    scanf("%d", &n);

    double **matriz = (double **)malloc(n * sizeof(double *));
    for (i = 0; i < n; i++)
        matriz[i] = (double *)malloc((n + 1) * sizeof(double));

    printf("Digite a matriz aumentada (linha por linha):\n");
    for (i = 0; i < n; i++)
        for (j = 0; j <= n; j++)
            scanf("%lf", &matriz[i][j]);

    double *solucao = (double *)malloc(n * sizeof(double));

    eliminacaoGauss(n, matriz, solucao);

    printf("\nSolucao:\n");
    for (i = 0; i < n; i++)
        printf("x[%d] = %.6f\n", i, solucao[i]);

    for (i = 0; i < n; i++)
        free(matriz[i]);
    free(matriz);
    free(solucao);

    return 0;
}