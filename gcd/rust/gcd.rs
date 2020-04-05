fn main() {
    let m = 18;
    let n = 30;
    println!("{} と {} の最大公約数は {}", m, n, gcd(m, n));
}

fn gcd(m: i32, n: i32) -> i32 {
    if n == 0 {
        return m;
    }
    return gcd(n, m % n);
}
