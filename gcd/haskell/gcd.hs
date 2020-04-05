thegcd :: Integer -> Integer -> Integer
thegcd m 0 = m
thegcd m n = thegcd n (mod m n)

main = let m = 18
           n = 30
        in putStrLn (show m ++ "と" ++ show n ++ "の最大公約数は" ++ show (thegcd m n))
