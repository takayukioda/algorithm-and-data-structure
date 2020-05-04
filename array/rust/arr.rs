const SIZE: i32 = 5
fn main() {
	let arr: [i32; SIZE] = [1, 2, 3, 4, 5]
}

fn insert(mut arr: [T; SIZE], idx: i32, val: T) -> [T; SIZE] {
    if idx >= SIZE || idx < 0 {
    }
}

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
