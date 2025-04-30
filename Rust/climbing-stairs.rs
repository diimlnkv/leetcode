//Time:
//Space:

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        if n < 3 {
            return n
        }
        let mut dp = vec![0; (n + 1) as usize];

        dp[1] = 1;
        dp[2] = 2;

        for i in 3..=n {
            dp[i as usize] = dp[(i-1) as usize] + dp[(i-2) as usize]
        }
        return dp[n as usize]
    }
}