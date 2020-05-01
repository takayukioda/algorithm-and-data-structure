theinsert :: [a] -> Int -> a -> [a]
theinsert x i v
  | i >= length x = x
  | i < 0 = x
  | otherwise = theinsert' x i v
    where 
        theinsert' :: [a] -> Int -> a -> [a]
        theinsert' x 0 v = [v] ++ (tail x)
        theinsert' (x:xs) i v = [x] ++ theinsert' xs (i - 1) v

thedelete :: [a] -> Int -> [a]
thedelete x i
  | i >= length x = x
  | i < 0 = x
  | otherwise = thedelete' x i
    where
        thedelete' :: [a] -> Int -> [a]
        thedelete' x 0 = tail x
        thedelete' (x:xs) i = [x] ++ thedelete' xs (i - 1)

main = let arr = [1, 2, 3, 4, 5]
           idx = 2
           val = 20
        in putStrLn (show (theinsert arr idx val) ++ show (thedelete arr idx))
