pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
   fn is_greater_or_equal_to(input: &[Vec<u64>], n: u64, row_index: usize) -> bool {
       for i in 0..input[0].len() {
           if n < input[row_index][i] {
               return false
           }
       }
       true
   }

   fn is_smaller_than_or_equal_to(input: &[Vec<u64>], n: u64, col_index: usize) -> bool {
       for i in 0..input.len() {
           if n > input[i][col_index] {
               return false
           }
       }
       true
   }

    let mut result = vec![];

    for row_index in 0..input.len() {
        for col_index in 0..input[0].len() {
            if is_greater_or_equal_to(input, input[row_index][col_index], row_index) && 
            is_smaller_than_or_equal_to(input, input[row_index][col_index], col_index) {
                result.push((row_index, col_index))
            }
        }
    }
    result
}
