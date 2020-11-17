package main

func canJump(nums []int) bool {
    if len(nums) == 1 {return true}
    q := []int{0}
    visited := make([]int, len(nums))
    for len(q) != 0 {
        count := len(q)
        for count > 0 {
            // fmt.Println("count ", count)
            count--
            
            top := q[0]
            for i:=1; i<=nums[top] && top + i < len(nums); i++ {
                newPos := top+i
                // fmt.Println(newPos)
                if visited[newPos] == 1 {continue}
                visited[newPos] = 1
                if newPos == len(nums)-1 {return true}
                q = append(q, newPos)
            }
            
            if len(q) > 1 {
                q = q[1:]
            } else {
                q = []int{}
            }
        }
    }
    return false
}