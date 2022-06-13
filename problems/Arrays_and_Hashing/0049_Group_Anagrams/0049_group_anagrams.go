package problems

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]byte][]string)
    for _, str := range strs {
        count := [26]byte{}
        for _, c := range str {
            count[c - 'a'] += 1
        }
        anagrams[count] = append(anagrams[count], str)
    }

    res := [][]string{}
    for _, v := range anagrams {
        res = append(res, v)
    }
    return res
}
