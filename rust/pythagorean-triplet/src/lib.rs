use std::collections::HashSet;

pub fn find(sum: u32) -> HashSet<[u32; 3]> {
	let mut result = HashSet::new();
	for element1 in 0..sum/3 {
		for element2 in element1+1..sum/2 {
			let element3 = sum - element1 - element2;
			if element1.pow(2) + element2.pow(2) == element3.pow(2) {
				result.insert([element1, element2, element3]);		
			}
		}
	}
	result
}
