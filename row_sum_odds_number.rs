fn row_sum_odd_numbers(n: i64) -> i64 {
    let first: i64 = 1 + n * (n - 1);
    let sum: i64;
    sum = first * n + n * (n - 1);
    return sum;
}

#[test]
fn returns_expected() {
    assert_eq!(row_sum_odd_numbers(1), 1);
    assert_eq!(row_sum_odd_numbers(42), 74088);
}

// fn main() {
//     let sum = row_sum_odd_numbers(4);
//     println!("res:{}", sum);
// }
