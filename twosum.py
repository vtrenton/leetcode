def twoSum(nums: list[int], target: int) -> list[int]:
   # Create a dictionary to store the numbers and their indices
   num_to_index = {}
   # Loop through the list of numbers
   for i, num in enumerate(nums):
       # Calculate the complement of the current number
       complement = target - num
       # Check if the complement exists in the dictionary
       if complement in num_to_index:
           # If it exists, return the indices of the complement and the current number
           return [num_to_index[complement], i]
       # Otherwise, add the current number and its index to the dictionary
       num_to_index[num] = i
   # If no pair is found, return an empty list
   return []

# Example usage
nums = [1, 8, 4, 15, 3, 6]
target = 18
print(twoSum(nums, target))


  #for i in range(len(nums)):
    #for j in range(len(nums)):
      #if j == i:
          #continue
      #result = nums[i] + nums[j]
      #if result == target:
        #return [i, j]

#twoSum(nums, target)
