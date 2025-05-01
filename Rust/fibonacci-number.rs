//Time:
//Space:

impl Solution {
    pub fn fib(n: i32) -> i32 {
        let mut pd = vec![0; (n+1) as usize];
        if n < 2 {
            return n
            }

        pd[0] = 0;
        pd[1] = 1;

        for i in 2..=n {
            pd[i as usize] = pd[(i-1) as usize] + pd[(i-2) as usize]
        }

        return pd[n as usize] as i32
    }
}