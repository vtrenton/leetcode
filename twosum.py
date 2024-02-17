nums = [1,8,4,15,3,6]
target = 18
def twoSum(nums: list[int], target: int) -> list[int]:
  for i in range(len(nums)):
    m = map(nums[i], i)
    if nums[i] - target in m:
      print(f"winrar {i} {m[i]}")
    else:
      print("Nope")
  #for i in range(len(nums)):
    #for j in range(len(nums)):
      #if j == i:
          #continue
      #result = nums[i] + nums[j]
      #if result == target:
        #return [i, j]

twoSum(nums, target)
