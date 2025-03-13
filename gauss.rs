use std::error::Error;
use std::io;
use std::io::Write;

type Matrix = Vec<Vec<f64>>;

fn eliminacao_gauss(matriz: &mut Matrix) -> Result<Vec<f64>, Box<dyn Error>> {
    let n = matriz.len();
    let mut solucao = vec![0.0; n];

    // Fase de eliminação progressiva
    for i in 0..n {
        // Pivotamento parcial
        let mut max_linha = i;
        for j in i..n {
            if matriz[j][i].abs() > matriz[max_linha][i].abs() {
                max_linha = j;
            }
        }
        matriz.swap(i, max_linha);

        // Verifica matriz singular
        if matriz[i][i].abs() < 1e-10 {
            return Err("Matriz singular, sem solução única".into());
        }

        // Eliminação
        for j in (i + 1)..n {
            let fator = matriz[j][i] / matriz[i][i];
            for k in i..(n + 1) {
                matriz[j][k] -= fator * matriz[i][k];
            }
        }
    }

    // Substituição regressiva
    for i in (0..n).rev() {
        solucao[i] = matriz[i][n];
        for j in (i + 1)..n {
            solucao[i] -= matriz[i][j] * solucao[j];
        }
        solucao[i] /= matriz[i][i];
    }

    Ok(solucao)
}

fn main() -> Result<(), Box<dyn Error>> {
    print!("Número de variáveis: ");
    io::stdout().flush()?;
    
    let mut input = String::new();
    io::stdin().read_line(&mut input)?;
    let n: usize = input.trim().parse()?;
    
    let mut matriz: Matrix = Vec::with_capacity(n);
    println!("Digite a matriz aumentada (linha por linha):");

    for _ in 0..n {
        input.clear();
        io::stdin().read_line(&mut input)?;
        let linha: Vec<f64> = input
            .split_whitespace()
            .map(|x| x.parse().unwrap_or_else(|_| {
                eprintln!("Erro: Valor não numérico detectado");
                std::process::exit(1);
            }))
            .collect();
        
        if linha.len() != n + 1 {
            eprintln!("Erro: Número inválido de elementos na linha");
            std::process::exit(1);
        }
        matriz.push(linha);
    }

    let mut matriz_copy = matriz.clone();
    let solucao = eliminacao_gauss(&mut matriz_copy)?;

    println!("\nSolução:");
    for (i, valor) in solucao.iter().enumerate() {
        println!("x[{}] = {:.6}", i, valor);
    }

    Ok(())
}