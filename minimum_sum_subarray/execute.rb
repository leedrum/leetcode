def minimum_sum_subarray(nums, l, r)
    # find the list sub-array with the condition l <= SIZE <= r

    min = -1

    #input [5,8,-6,7,9], 1,3
        # [5]
        # [8]
        # [5,8],
        # [-6], [8,-6], [5,8,-6]
        # [7], [-6,7], [8,-6,7]
        # [9], [7,9], [-6,7,9]
    (0..nums.size - 1).each do |i|
        temp_sum = 0
        count_item = 0
        (0..r).each do |k|
            next if i - k < 0

            if !nums[i-k].nil? && count_item <= r
                temp_sum += nums[i-k]
                count_item += 1
                min = temp_sum if count_item.between?(l,r) && temp_sum > 0 && (min < 0 || temp_sum < min)
            end
        end
    end

    return min
end

p "result: #{minimum_sum_subarray([3,-2,1,4],2,3) == 1}"
p "result: #{minimum_sum_subarray([-2,2,-3,1],2,3) == -1}"
p "result: #{minimum_sum_subarray([1,2,3,4],2,3) == 3}"
p "result: #{minimum_sum_subarray([-12,8],1,1) == 8}"
p "result: #{minimum_sum_subarray([5,8,-6],1,3) == 2}"
p "result: #{minimum_sum_subarray([7,3],2,2) == 10}"
p "result: #{minimum_sum_subarray([9,11,-9],3,3) == 11}"
p "result: #{minimum_sum_subarray([25,-9],1,1) == 25}"
