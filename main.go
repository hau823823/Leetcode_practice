package main

import (
	"fmt"
	"leetcode/longestCommonSubsequence"
)

/**
                                                     _ooOoo
                                                    o8888888o
                                                    88" . "88
                                                    (| -_- |)
                                                    O\  =  /O
                                                  ___/`---'\____
                                               .'  \\|     |//  `.
                                              /  \\|||  :  |||//  \
                                             /  _||||| -:- |||||_  \
                                             |   | \\\  -  /// |   |
                                             | \_|  ''\---/''  |   |
                                             \  .-\__       __/-.  /
                                           ___`. .'  /--.--\ `. . __
                                        ."" '<  `.___\_<|>_/__.'  >'"".
                                       | | :  `- \`.;`\ _ /`;.`/ - ` : | |
                                       \  \ `-.   \_ __\ /__ _/   .-` /  /
                                  ======`-.____`-.___\_____/___.-`____.-'======
                                                     `=---='
                               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                                              佛祖保佑       永無BUG
**/

func main() {

	// climbing stairs
	/**
	for i := 1; i <= 45; i++ {
		fmt.Printf("input: %d, steps: %d \n", i, climbingStairs.ClimbStairs(i))
	}
	*/

	// best time to buy and sell stock
	/**
	prices1 := []int{7,1,5,3,6,4}
	fmt.Println(bestTimeBuySellStock1.MaxProfitDP(prices1))
	prices2 := []int{7,6,4,3,1}
	fmt.Println(bestTimeBuySellStock1.MaxProfitDP((prices2)))
	*/

	// best time to buy and sell stock 2
	/**
	prices := []int{7,1,5,3,6,4}
	fmt.Println(bestTimeBullSellStocks2.MaxProfit(prices))

	prices1 := []int{1,2,3,4,5}
	fmt.Println(bestTimeBullSellStocks2.MaxProfit(prices1))
	*/

	// maximum subarray
	/**
	nums1 := []int{5,4,-1,7,8}
	fmt.Println(maximumSubarray.MaxSubArraryGreedy(nums1))
	*/

	// house robber
	/**
	nums1 := []int{1,2,3,1}
	fmt.Println(houseRobber.Rob(nums1))
	nums2 := []int{2,1,1,2}
	fmt.Println(houseRobber.Rob(nums2))
	*/

	// house robber II
	/**
	nums := []int{2, 3, 2}
	fmt.Println(houseRobberII.Rob(nums))
	*/

	// string to integer (atoi)
	//str1 := "   -42"
	//str2 := "word and 987"
	//str3 := "+-3"
	//str4 := "00000-42a1234"
	//str5 := "20000000000000000000"
	//fmt.Println(string2Integer.MyAtoi(str5))

	// implement strStr()
	// same as question "find first occurence chat in strings"
	/**
	haystack := "sadbutsad";
	needle := "sad"
	fmt.Println(implementStr.StrStrUseStrings(haystack, needle))

	haystack1 := "leetcode";
	needle1 := "leeto"
	fmt.Println(implementStr.StrStrUseStrings(haystack1, needle1))

	haystack2 := "sssss";
	needle2 := "ss"
	fmt.Println(implementStr.StrStrUseStrings(haystack2, needle2))
	*/

	// two sum
	/**
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum.TwoSum(nums, target))

	nums1 := []int{3,2,4}
	target1 := 6
	fmt.Println(twoSum.TwoSum(nums1, target1))

	nums2 := []int{3,3}
	target2 := 6
	fmt.Println(twoSum.TwoSum(nums2, target2))
	*/

	// Remove Duplicates from Sorted Array
	/**
	nums := []int{1,1,2}
	nums1 := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicated.RemoveDuplicates(nums), nums)
	fmt.Println(removeDuplicated.RemoveDuplicates(nums1), nums1)
	*/

	// Convert Sorted Array to Binary Search Tree
	/**
	nums := []int{-10,-3,0,5,9}
	bst := convertArray2BST.SortedArrayToBST(nums)
	result := binaryTreeLevelOrderTraversal.LevelOrderLoop(bst)
	fmt.Println(result)
	*/

	// fizz buzz
	/**
	n := 3
	fmt.Println(fizzBuzz.FizzBuzzBrute(n))

	n1 := 5
	fmt.Println(fizzBuzz.FizzBuzzBrute(n1))

	n2 := 15
	fmt.Println(fizzBuzz.FizzBuzzBrute(n2))
	*/

	// count primes
	//n := 10
	//fmt.Println(countPrimes.CountPrimes(n))

	// power of three
	//n := 1
	//fmt.Println(powerOfThree.IsPowerOfThree(n))

	// roman to interger
	/**
	s := "VIII"
	fmt.Println(romanToInteger.RomanToInt(s))

	s1 := "LVIII"
	fmt.Println(romanToInteger.RomanToInt(s1))

	s2 := "MCMXCIV"
	fmt.Println(romanToInteger.RomanToInt(s2))
	*/

	// number of 1 bits
	/**
	n := uint32(4294967293) // 11111111111111111111111111111101
	//n := uint32(128) // 00000000000000000000000010000000
	fmt.Println("num(decimal): ", n)
	fmt.Printf("num(binary): %b\n", n)
	fmt.Println("\nANS: ",numberOf1Bits.HammingWeight(n))
	*/

	// hamming distance
	/**
	x := 3
	y := 1
	fmt.Println(hammingDistance.HammingDistance(x, y))
	*/

	// reverse bits
	/**
	n := uint32(43261596)
	fmt.Printf("n(decimal): %d\n", n)
	fmt.Printf("n(binary): %32b\n", n)

	ans := reverseBits.ReverseBits(n)
	fmt.Printf("Ans(decimal): %d\n", ans)
	fmt.Printf("Ans(binary): %32b\n", ans)
	*/

	// pascal's triangle
	/**
	numRows := 1
	fmt.Println(pascalsTriangle.Generate(numRows))
	*/

	// valid parentheses
	//s := "([]"
	//fmt.Println(validParentheses.IsValid(s))

	//s1 := "(){}}{"
	//fmt.Println(validParentheses.IsValid(s1))

	// missing number
	/**
	nums := []int{3,0,1}
	fmt.Println(missingNumber.MissingNumber(nums))
	*/

	// shuffle an array
	/**
	p := shuffleArrary.Constructor([]int{1,2,3})
	for i := 0; i < 10; i ++ {
		fmt.Println("Shuffle: ", p.Shuffle())
		fmt.Println("Reset: ", p.Reset())
	}
	*/

	// min stack
	/**
	obj := minStack.Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Push(-1)
	obj.Push(-5)
	obj.Push(5)

	fmt.Println(obj)
	fmt.Println("GetMin: ", obj.GetMin())
	fmt.Println("Top: ", obj.Top())
	fmt.Println("---POP---")
	obj.Pop()
	fmt.Println(obj)
	fmt.Println("GetMin: ", obj.GetMin())
	fmt.Println("Top: ", obj.Top())
	*/

	//three sum
	/**
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum.ThreeSum2PT(nums))
	*/

	// set matrix zero
	/**
	matrix := [][]int{
		{1,1,1},
		{1,0,1},
		{1,1,1},
	}
	setMatrixZero.SetZeroesOp(matrix)
	fmt.Println(matrix)

	matrix1 := [][]int{
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}
	setMatrixZero.SetZeroesOp(matrix1)
	fmt.Println(matrix1)
	*/

	// remove element
	/**
	nums := []int{3,2,2,3}
	val := 3
	fmt.Println(removeElement.RemoveElement(nums, val))
	fmt.Println(nums)

	nums1 := []int{0,1,2,2,3,0,4,2}
	val1 := 2
	fmt.Println(removeElement.RemoveElement(nums1, val1))
	fmt.Println(nums1)

	nums2 := []int{4,5}
	val2 := 5
	fmt.Println(removeElement.RemoveElement(nums2, val2))
	fmt.Println(nums2)
	*/

	// zigzag
	/**
	s := "PAYPALISHIRING"
	fmt.Println(zigZagConversation.Convert(s, 3))
	*/

	// group anagrams
	/**
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams.GroupAnagramsOp(strs))

	strs1 := []string{""}
	fmt.Println(groupAnagrams.GroupAnagramsOp(strs1))

	strs2 := []string{"a"}
	fmt.Println(groupAnagrams.GroupAnagramsOp(strs2))
	*/

	// rotate array
	/**
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotateArray.RotateOP1(nums, k)
	fmt.Println(nums)

	nums1 := []int{-1, -100, 3, 99}
	k1 := 2
	rotateArray.RotateOP1(nums1, k1)
	fmt.Println(nums1)
	*/

	// plus one
	/**
	digits := []int{1,2,9}
	fmt.Println(plusOne.PlusOne2(digits))

	digits1 := []int{9}
	fmt.Println(plusOne.PlusOne2(digits1))
	*/

	// sqrt(x)
	/**
	x := 4
	fmt.Println(sqrt.MySqrtBS(x))

	x1 := 8
	fmt.Println(sqrt.MySqrtBS(x1))
	*/

	// contains duplicate
	/**
	nums := []int{1,2,3,1}
	fmt.Println(containsDuplicate.ContainsDuplicate(nums))

	nums1 := []int{1,2,3,4}
	fmt.Println(containsDuplicate.ContainsDuplicate(nums1))
	*/

	// single number
	/**
	nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber.SingleNumberOP(nums))

	nums1 := []int{2,2,1}
	fmt.Println(singleNumber.SingleNumberOP(nums1))
	*/

	// longest substring without repeating characters
	/**
	s := "abcabcbb"
	fmt.Println(longestSubStrWithoutRptChts.LengthOfLongestSubStringOP(s))

	s1 := "bbbbb"
	fmt.Println(longestSubStrWithoutRptChts.LengthOfLongestSubStringOP(s1))

	s2 := "pwwkew"
	fmt.Println(longestSubStrWithoutRptChts.LengthOfLongestSubStringOP(s2))

	s3 := "abba"
	fmt.Println(longestSubStrWithoutRptChts.LengthOfLongestSubStringOP(s3))
	*/

	// majority element
	/**
	nums := []int{3,2,3}
	fmt.Println(majorityElement.MajorityElementOp(nums))

	nums1 := []int{2,2,1,1,1,2,2,3,3,4}
	fmt.Println(majorityElement.MajorityElementOp(nums1))
	*/

	// Longest Palindromic Substring
	/**
	s := "babad"
	fmt.Println(longestPalindromicSubstring.LongestPalindromeDP(s))

	s1 := "cbbd"
	fmt.Println(longestPalindromicSubstring.LongestPalindromeDP(s1))

	s2 := "bbb"
	fmt.Println(longestPalindromicSubstring.LongestPalindromeDP(s2))
	*/

	// Longest Consecutive Sequence
	/**
	nums := []int{100,4,200,1,3,2}
	fmt.Println(longestConsecutiveSequence.LongestConsecutiveOP(nums))

	nums1 := []int{0,3,7,2,5,8,4,6,0,1}
	fmt.Println(longestConsecutiveSequence.LongestConsecutiveOP(nums1))
	*/

	// length of last word
	//s := "Hello World"
	//fmt.Println(lengthOfLastWord.LengthOfLastWord(s))

	// ransom note
	/**
	r := "aa"
	m := "ab"
	fmt.Println(ransomNote.CanConstructOp(r, m))

	r1 := "aa"
	m1 := "aab"
	fmt.Println(ransomNote.CanConstructOp(r1, m1))
	*/

	// my pow
	//start := time.Now()
	//x := float64(2)
	//n := 10
	//elapsed := time.Since(start)
	//fmt.Println(pow.MyPowOp(x, n))
	//fmt.Println("time: ", elapsed)

	//x1 := float64(2)
	//n1 := 7
	//fmt.Println(pow.MyPowOp(x1, n1))

	// increasing tirplet subsequence
	/**
	nums := []int{1,2,3,4,5}
	fmt.Println(increasingTripletSubSequence.IncreasingTriplet(nums))

	nums1 := []int{2,1,5,0,4,6}
	fmt.Println(increasingTripletSubSequence.IncreasingTriplet(nums1))
	*/

	// count and say
	/**
	test := "223314444411"
	fmt.Println(countNsay.Helper1(test))
	fmt.Println(countNsay.Helper2(countNsay.Helper1(test)))
	fmt.Println(countNsay.CountAndSay(5))
	*/

	// move zeroes
	/**
	nums := []int{4,2,4,0,0,3,0,5,1,0}
	moveZeroes.MoveZeroes(nums)
	fmt.Println(nums)

	nums1 := []int{1,1}
	moveZeroes.MoveZeroes(nums1)
	fmt.Println(nums1)
	*/

	// valid sudoku
	/**
	board := [][]byte {
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	fmt.Println(validSudoku.IsValidSudoku(board))
	*/

	// rotate image
	/**
	matrix := [][]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	rotateImage.Rotate(matrix)
	fmt.Println(matrix)

	matrix1 := [][]int {
		{5,1,9,11},
		{2,4,8,10},
		{13,3,6,7},
		{15,14,12,16},
	}
	rotateImage.Rotate(matrix1)
	fmt.Println(matrix1)
	*/

	// add two numbers
	/**
	l1 := node.Ints2ListNode([]int{9,9,9,9,9,9,9})
	l2 := node.Ints2ListNode([]int{9,9,9,9})
	ans := add2Numbers.AddTwoNumbers(l1, l2)
	fmt.Println(node.ListNode2Ints(ans))

	l1_1 := node.Ints2ListNode([]int{2,4,3})
	l2_1 := node.Ints2ListNode([]int{5,6,4})
	ans1 := add2Numbers.AddTwoNumbers(l1_1, l2_1)
	fmt.Println(node.ListNode2Ints(ans1))
	*/

	// reverse string
	/**
	s := []byte{'h','e','l','l','o'}
	reverseString.ReverseString(s)
	fmt.Println(string(s))
	*/

	// reverse int
	/**
	x := 123
	fmt.Println(reverseInt.Reverse(x))

	x1 := -120
	fmt.Println(reverseInt.Reverse(x1))
	*/

	// find first occurence chat in strings
	/**
	haystack := "butsad"
	needle := "sad"
	fmt.Println(find1StOccurenceInStr.StrStr(haystack, needle))

	haystack1 := "leetcode"
	needle1 := "leet0"
	fmt.Println(find1StOccurenceInStr.StrStr(haystack1, needle1))

	haystack2 := "a"
	needle2 := "a"
	fmt.Println(find1StOccurenceInStr.StrStr(haystack2, needle2))
	*/

	// jump game
	/**
	nums := []int{2,3,1,1,4}
	fmt.Println(jumpGame.CanJump(nums))

	nums1 := []int{3,2,1,0,4}
	fmt.Println(jumpGame.CanJump(nums1))
	*/

	// odd even linked list
	/**
	head := node.Ints2ListNode([]int{1,2,3,4,5})
	ans := node.ListNode2Ints(oddEvenLinkedList.OddEvenList(head))
	fmt.Println(ans)
	*/

	// longest common prefix
	/**
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix.LongestCommonPrefix(strs))
	fmt.Printf("\n")

	strs1 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix.LongestCommonPrefix(strs1))
	fmt.Printf("\n")

	strs2 := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix.LongestCommonPrefix(strs2))
	fmt.Printf("\n")
	*/

	// reverse words in a string
	/**
	s := "  hello world  "
	fmt.Println(s)
	fmt.Println(reverseWords.ReverseWords(s))
	*/

	// find first unique character in string
	/**
	s := "leetcode"
	fmt.Println(firstUniqueChrInStr.FirstUniqChar(s))

	s1 := "loveleetcode"
	fmt.Println(firstUniqueChrInStr.FirstUniqChar(s1))
	*/

	// valid anagram
	/**
	s := "anagram"
	t := "nagaram"
	fmt.Println(validAnagram.IsAnagram(s, t))

	s1 := "aacc"
	t1 := "ccac"
	fmt.Println(validAnagram.IsAnagram(s1, t1))
	*/

	// valid palindrome
	/**
	s := "A man, a plan, a canal: Panama"
	fmt.Println(validPalindrome.IsPalindrome(s))

	s1 := "race a car"
	fmt.Println(validPalindrome.IsPalindrome(s1))

	s2 := "0p"
	fmt.Println(validPalindrome.IsPalindrome(s2))
	*/

	// intersection of two linked list
	/**
	listA := []int{4, 1, 8, 4, 5}
	listB := []int{5, 6, 1, 8, 4, 5}
	skipA := 2
	skipB := 3
	headA, headB := intersection2LinkedList.Generate(listA, listB, skipA, skipB)
	fmt.Println(intersection2LinkedList.GetIntersectionNodeOP(headA, headB))
	*/

	/**
	listA_1 := []int{2, 6, 4}
	listB_1 := []int{1, 5}
	skipA_1 := 3
	skipB_1 := 2
	headA_1, headB_1 := intersection2LinkedList.Generate(listA_1, listB_1, skipA_1, skipB_1)
	fmt.Println(intersection2LinkedList.GetIntersectionNodeOP(headA_1, headB_1))
	*/

	/**
	listA_2 := []int{1, 9, 1, 2, 4}
	listB_2 := []int{1, 2, 4}
	skipA_2 := 3
	skipB_2 := 1
	headA_2, headB_2 := intersection2LinkedList.Generate(listA_2, listB_2, skipA_2, skipB_2)
	fmt.Println(intersection2LinkedList.GetIntersectionNodeOP(headA_2, headB_2))
	*/

	// word pattern
	/**
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern.WordPattern(pattern, s))

	pattern1 := "abba"
	s1 := "dog cat cat fish"
	fmt.Println(wordPattern.WordPattern(pattern1, s1))

	pattern2 := "abba"
	s2 := "dog dog dog dog"
	fmt.Println(wordPattern.WordPattern(pattern2, s2))
	*/

	// binary tree inoder traversal
	/**
	ints := []int{4, 2, 6, 1, 3, 5, 7}
	root := node.Ints2TreeNode(ints)
	ans := node.Tree2PostOrderStack(root)
	fmt.Println(ans)
	*/

	// maximum depth of binary tree
	/**
	ints := []int{3, 9, 20, node.NULL, node.NULL, 15, 7}
	root := node.Ints2TreeNode(ints)
	fmt.Println(maximumDepthBT.MaxDepthTEST(root, 0))
	*/

	// minimum depth of binary tree
	/**
	ints := []int{3, 9, 20, node.NULL, node.NULL, 15, 7}
	root := node.Ints2TreeNode(ints)
	fmt.Println(minDepthOfBinaryTree.MinDepth(root))
	*/

	// is subsequence
	/**
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubquence.IsSubsequence(s, t))

	s1 := "axc"
	t1 := "ahbgdc"
	fmt.Println(isSubquence.IsSubsequence(s1, t1))
	*/

	// product expect self
	/**
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf.ProductExceptSelfOp(nums))
	fmt.Printf("\n")

	nums1 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf.ProductExceptSelfOp(nums1))
	fmt.Printf("\n")
	*/

	// delete node in a link list
	//head := node.Ints2ListNode([]int{4,5,1,9})

	// remove nth node from end of list
	/**
	head := node.Ints2ListNode([]int{1,2,3,4,5})
	n := 2
	rmNthNodeFromEndOfLinkedList.RemoveNthFromEnd2PT(head, n)
	fmt.Println(node.ListNode2Ints(head))
	*/

	// reverse linked list
	/**
	head := node.Ints2ListNode([]int{1,2,3,4,5})
	ans := node.ListNode2Ints(reverseLinkedList.ReverseList(head))
	fmt.Println(ans)
	*/

	// happy number
	/**
	n := 19
	fmt.Println(happyNumber.IsHappy(n))

	n1 := 2
	fmt.Println(happyNumber.IsHappy(n1))
	*/

	// two sum input array is sorted
	/**
	numbers := []int{2,3,4}
	target := 6
	fmt.Println(twoSumII.TwoSumOP(numbers, target))
	*/

	// merge two sorted lists
	/**
	list1 := node.Ints2ListNode([]int{})
	list2 := node.Ints2ListNode([]int{0})
	ans := node.ListNode2Ints(merge2SortedLists.MergeTwoLists(list1, list2))
	fmt.Println(ans)
	*/

	// Binary Tree Level Order Traversal
	/**
	ints := []int{3, 9, 20, node.NULL, node.NULL, 15, 7}
	root := node.Ints2TreeNode(ints)
	fmt.Println(binaryTreeLevelOrderTraversal.LevelOrderLoop(root))

	ints1 := []int{1}
	root1 := node.Ints2TreeNode(ints1)
	fmt.Println(binaryTreeLevelOrderTraversal.LevelOrderLoop(root1))

	ints2 := []int{}
	root2 := node.Ints2TreeNode(ints2)
	fmt.Println(binaryTreeLevelOrderTraversal.LevelOrderLoop(root2))
	*/

	// Binary Tree Zigzag Level Order Traversal
	/**
	ints := []int{3, 9, 20, node.NULL, node.NULL, 15, 7}
	root := node.Ints2TreeNode(ints)
	fmt.Println(binaryTreeZigzagLevelOrderTraversal.ZigZagLevelOrderLoop(root))

	ints1 := []int{1, 2, 3, 4, node.NULL, node.NULL, 5}
	root1 := node.Ints2TreeNode(ints1)
	fmt.Println(binaryTreeZigzagLevelOrderTraversal.ZigZagLevelOrderLoop(root1))
	*/

	// construct binary tree from preorder and inorder traversal
	/**
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	tree := constructBinaryTreeFromPreOrderAndInOrder.BuildTreeOP(preorder, inorder)
	fmt.Println("preorder: ", node.Tree2PreOrder(tree))
	fmt.Println("inorder: ", node.Tree2InOrder(tree))
	*/

	// palindrome linked list
	/**
	head := node.Ints2ListNode([]int{1, 2, 2, 1})
	fmt.Println(palindromeLinkedList.IsPalindromeTP(head
	*/

	// contains duplicate II
	/**
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsDuplicateII.ContainsNearbyDuplicate(nums, k))

	nums1 := []int{1,0,1,1}
	k1 := 1
	fmt.Println(containsDuplicateII.ContainsNearbyDuplicate(nums1, k1))

	nums2 := []int{1,2,3,1,2,3}
	k2 := 2
	fmt.Println(containsDuplicateII.ContainsNearbyDuplicate(nums2, k2))
	*/

	// linked list cycle
	/**
	head := node.Ints2ListNode([]int{3,2,0,-4})
	fmt.Println(linkedListCycle.HasCycle(head))
	*/

	// validate binary search tree
	/**
	root := node.Ints2TreeNode([]int{2,1,3})
	fmt.Println(validateBST.IsValidBST(root))

	root1 := node.Ints2TreeNode([]int{5,1,4,node.NULL,node.NULL,3,6})
	fmt.Println(validateBST.IsValidBST(root1))
	*/

	// symmetric tree
	/**
	root := node.Ints2TreeNode([]int{1,2,2,3,4,4,3})
	fmt.Println(symmetricTree.IsSymmetricRecursive(root))

	root1 := node.Ints2TreeNode([]int{1,2,2,node.NULL,3,node.NULL,3})
	fmt.Println(symmetricTree.IsSymmetricRecursive(root1))
	*/

	// minimum aboslute difference in bst
	/**
	root := node.Ints2TreeNode([]int{4,2,6,1,3})
	fmt.Println(minAbsDifInBFS.GetMinimumDifference(root))
	*/

	// kth smallest element in bst
	/**
	root := node.Ints2TreeNode([]int{5,3,6,2,4,node.NULL,node.NULL,1})
	k := 3
	fmt.Println(kthSmallestElementInBST.KthSmallest(root, k))
	fmt.Println(kthSmallestElementInBST.KthSmallest2(root, k))
	fmt.Println(kthSmallestElementInBST.KthSmallestCount(root, k))
	fmt.Println(kthSmallestElementInBST.KthSmallestOP(root, k))
	*/

	// numbers of islands
	/**
	grid := [][]byte {
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	fmt.Println(numberOfIslands.NumIslandsBFS(grid))
	grid = [][]byte {
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	fmt.Println(numberOfIslands.NumIslandsDFS(grid))
	*/

	// generate parenthesis
	/**
	n := 3
	fmt.Println(generateParenthesis.GenerateParenthesis(n))
	*/

	// letter of combinations of a phone number
	/**
	digits := "234"
	fmt.Println(letterCombinationsOfPhoneNumber.LetterCombinations(digits))
	*/

	// merge sorted arrary
	/**
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	mergeSortedArrary.Merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	*/

	// first bad version
	/**
	firstBadVersion.Bad = 1
	fmt.Println(firstBadVersion.FirstBadVersion(2))
	*/

	// permutations
	/**
	nums := []int{1,2,3}
	fmt.Println(permutations.Permute(nums))
	*/

	// subsets
	/**
	nums := []int{1, 2, 2}
	fmt.Println(subset.SubsetsBackTrack(nums))
	*/

	// word search
	/**
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(wordSearch.Exist(board, word))

	word2 := "SEE"
	fmt.Println(wordSearch.Exist(board, word2))

	word3 := "ABCB"
	fmt.Println(wordSearch.Exist(board, word3))
	*/

	// bubble sort
	/**
	nums := []int{2,0,2,1,1,0}
	sortColors.SortColorsTP(nums)
	fmt.Println(nums)
	*/

	// top k frequent elements
	/**
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequentElements.TopKFrequentOp(nums, k))

	nums1 := []int{3,0,1,0}
	k1 := 2
	fmt.Println(topKFrequentElements.TopKFrequent(nums1, k1))
	*/

	// find the kth largest elements in an arrary
	/**
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(kthLargestElementInArray.FindKthLargest2(nums, k))

	nums1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k1 := 4
	fmt.Println(kthLargestElementInArray.FindKthLargest2(nums1, k1))
	*/

	// find peak element
	/**
	nums := []int{1,2,3,1}
	fmt.Println(findPeakElement.FindPeakElementOp(nums))

	nums1 := []int{1,2,1,3,5,6,4}
	fmt.Println(findPeakElement.FindPeakElementOp(nums1))
	*/

	// search for a ragne
	/**
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchForRange.SearchRange(nums, target))

	nums1 := []int{1}
	target1 := 1
	fmt.Println(searchForRange.SearchRange(nums1, target1))
	*/

	// merge intervals
	/**
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(mergeIntervals.MergeOp(intervals))

	intervals1 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(mergeIntervals.MergeOp(intervals1))
	*/

	// search in rotated sorted array
	/**
	nums := []int{6, 7, 0, 1, 2, 4, 5}
	target := 0
	fmt.Println(searchInRotatedSortedArray.SearchOp(nums, target))
	*/

	// search a sd matrix II
	/**
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	fmt.Println(search2dMatrixII.SearchMatrix(matrix, target))
	*/

	// unique paths
	/**
	m := 3
	n := 7
	fmt.Println(uniquePath.UniquePaths(m, n))

	m1 := 3
	n1 := 2
	fmt.Println(uniquePath.UniquePaths(m1, n1))
	*/

	// coin change
	/**
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange.CoinChange(coins, amount))

	coins1 := []int{1}
	amount1 := 0
	fmt.Println(coinChange.CoinChange(coins1, amount1))
	*/

	// longest increasing subsequence
	/**
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(longestIncreasingSubsequence.LengthOfLISOp(nums))
	*/

	// factorial trailing zeroes
	/**
	n := 10
	fmt.Println(factorialTrailingZeroes.TrailingZeroes(n))
	*/

	// excel sheet column number
	/**
	columnTitle := "AAA"
	fmt.Println(excelSheetColumnNumber.TitleToNumberOp(columnTitle))
	*/

	// sum of two intergers
	/**
	a := 2
	b := 3
	fmt.Println(sumOf2Ints.GetSum(a, b))
	*/

	// Evaluate Reverse Polish Notation
	/**
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evaluateReversePolishNotation.EvalRPN(tokens))

	tokens1 := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	fmt.Println(evaluateReversePolishNotation.EvalRPN(tokens1))
	*/

	// task scheduler
	/**
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(taskScheduler.LeastIntervalPQ(tasks, n))

	tasks1 := []byte{'A','C','A','B','D','B'}
	n1 := 1
	fmt.Println(taskScheduler.LeastIntervalPQ(tasks1, n1))

	tasks2 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n2 := 3
	fmt.Println(taskScheduler.LeastIntervalPQ(tasks2, n2))
	*/

	// divide two integers
	/**
	dividend := 10
	divisor := 3
	fmt.Println(divide2Ints.Divide2(dividend, divisor))
	*/

	// fraction to recurring decimal
	/**
	numerator := 500
	denominator := 10
	fmt.Println(fraction2RecurringDecimal.FractionToDecimal(numerator, denominator))

	numerator1 := 4
	denominator1 := -333
	fmt.Println(fraction2RecurringDecimal.FractionToDecimal(numerator1, denominator1))
	*/

	// container with most water
	/**
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(containerWithMostWater.MaxArea(height))

	height1 := []int{1,1}
	fmt.Println(containerWithMostWater.MaxArea(height1))
	*/

	// longest repeating characters replacement
	/**
	s := "ABAB"
	k := 2
	fmt.Println(longestRepeatingChrReplacement.CharacterReplacement(s, k))

	s1 := "AABABBA"
	k1 := 1
	fmt.Println(longestRepeatingChrReplacement.CharacterReplacement(s1, k1))
	*/

	// permutation in string
	/**
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(permutationInString.CheckInclusionOP(s1, s2))

	s1_1 := "adc"
	s2_1 := "dcda"
	fmt.Println(permutationInString.CheckInclusionOP(s1_1, s2_1))
	*/

	// lru cache
	/**
	obj := lruCache.Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	*/

	// daily temperatures
	/**
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures.DailyTemperatures(temperatures))
	*/

	// car fleet
	/**
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet.CarFleetOp(target, position, speed))
	*/

	// binary search
	/**
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(binarySearch.Search(nums, target))

	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 2
	fmt.Println(binarySearch.Search(nums1, target1))
	*/

	// search a 2D matrix
	/**
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	fmt.Println(search2DMatrix.SearchMatrixOp2(matrix, target))

	matrix1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target1 := 13
	fmt.Println(search2DMatrix.SearchMatrixOp2(matrix1, target1))

	matrix2 := [][]int{{1, 3}}
	target2 := 3
	fmt.Println(search2DMatrix.SearchMatrixOp2(matrix2, target2))
	*/

	// Koko eating bananas
	/**
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(kokoEatingBananas.MinEatingSpeed(piles, h))

	piles1 := []int{30, 11, 23, 4, 20}
	h1 := 5
	fmt.Println(kokoEatingBananas.MinEatingSpeed(piles1, h1))
	*/

	// find minimum in rotated sorted array
	/**
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMinInRotatedSortedArray.FindMin(nums))
	*/

	// time based key-value store
	/**
	obj := timeBasedKeyValueStore.Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	*/

	// reorder list
	/**
	head := node.Ints2ListNode([]int{1, 2, 3, 4})
	reorderList.ReorderList(head)
	fmt.Println(node.ListNode2Ints(head))

	head1 := node.Ints2ListNode([]int{1, 2, 3, 4, 5})
	reorderList.ReorderList(head1)
	fmt.Println(node.ListNode2Ints(head1))
	*/

	// copy list with random pointer

	// find the duplicate number
	/**
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicateNumber.FindDuplicateBs(nums))
	*/

	// invert binary tree
	/**
	root := node.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9})
	fmt.Println(node.Tree2PreOrder(invertBinaryTree.InvertTree(root)))
	*/

	// diameter of binary tree
	/**
	root := node.Ints2TreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println(diameterOfBinaryTree.DiameterOfBinaryTree(root))
	*/

	// balanced binary tree
	/**
	root := node.Ints2TreeNode([]int{3, 9, 20, node.NULL, node.NULL, 15, 7})
	fmt.Println(balancedBinaryTree.IsBalancedOp(root))

	root1 := node.Ints2TreeNode([]int{1, 2, 2, 3, 3, node.NULL, node.NULL, 4, 4})
	fmt.Println(balancedBinaryTree.IsBalancedOp(root1))
	*/

	// same tree
	/**
	p := node.Ints2TreeNode([]int{1, 2, 3})
	q := node.Ints2TreeNode([]int{1, 2, 3})
	fmt.Println(sameTree.IsSameTree(p, q))

	p1 := node.Ints2TreeNode([]int{1, 2})
	q1 := node.Ints2TreeNode([]int{1, node.NULL, 2})
	fmt.Println(sameTree.IsSameTree(p1, q1))
	*/

	// subtree of another tree
	/**
	root := node.Ints2TreeNode([]int{3, 4, 5, 1, 2})
	subRoot := node.Ints2TreeNode([]int{4, 1, 2})
	fmt.Println(subTreeOfAnotherTree.IsSubtree(root, subRoot))

	root1 := node.Ints2TreeNode([]int{3, 4, 5, 1, 2, node.NULL, node.NULL, node.NULL, node.NULL, 0})
	subRoot1 := node.Ints2TreeNode([]int{4, 1, 2})
	fmt.Println(subTreeOfAnotherTree.IsSubtree(root1, subRoot1))
	*/

	// lowest common ancestor of binary search tree
	/**
	root := node.Ints2TreeNode([]int{6, 2, 8, 0, 4, 7, 9, node.NULL, node.NULL, 3, 5})
	p := root.Left  // 2
	q := root.Right // 8
	fmt.Println(lowestCommonAncestorOfBST.LowestCommonAncestorRecursive(root, p, q))

	p1 := root.Left // 2
	q1 := root.Left.Right // 4
	fmt.Println(lowestCommonAncestorOfBST.LowestCommonAncestorRecursive(root, p1, q1))
	*/

	// Binary Tree Right Side View
	/**
	root := node.Ints2TreeNode([]int{1, 2, 3, node.NULL, 5, node.NULL, 4})
	fmt.Println(binaryTreeRightSideView.RightSideView(root))
	*/

	// good nodes in binary tree
	/**
	root := node.Ints2TreeNode([]int{3, 1, 4, 3, node.NULL, 1, 5})
	fmt.Println(countGoodNodesInBinaryTree.GoodNodes(root))

	root1 := node.Ints2TreeNode([]int{3, 3, node.NULL, 4, 2})
	fmt.Println(countGoodNodesInBinaryTree.GoodNodes(root1))
	*/

	// last stone weight
	/**
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight.LastStoneWeightSort(stones))
	*/

	// kth largest element in stream
	/**
	obj := kthLargestElementInStream.Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
	*/

	// k closest points to origin
	/**
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	fmt.Println(kClosestPoints2Origin.KClosestOP(points, k))
	*/

	// Design Twitter
	/**
	twitter := designTwitter.Constructor()
	twitter.PostTweet(1, 5)             // User 1 posts a new tweet (id = 5).
	fmt.Println(twitter.GetNewsFeed(1)) // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
	twitter.Follow(1, 2)                // User 1 follows user 2.
	twitter.PostTweet(2, 6)             // User 2 posts a new tweet (id = 6).
	fmt.Println(twitter.GetNewsFeed(1)) // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	twitter.Unfollow(1, 2)              // User 1 unfollows user 2.
	twitter.GetNewsFeed(1)              // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
	*/

	// combination Sum
	/**
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum.CombinationSum(candidates, target))

	candidates1 := []int{2,3,5}
	target1 := 8
	fmt.Println(combinationSum.CombinationSum(candidates1, target1))
	*/

	// subset II
	/**
	nums := []int{1, 3, 2}
	fmt.Println("subset:   ", subset.SubsetsBackTrack(nums))
	fmt.Println("subsetII: ", subsetsII.SubsetsWithDup(nums))
	*.

	// combination sum II
	/**
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum.CombinationSum(candidates, target))
	fmt.Println(combinationSumII.CombinationSum2(candidates, target))
	*/

	// palidrome partitioning
	/**
	s := "aaab"
	fmt.Println(palindromePartitioning.Partition(s))
	*/

	// implement trie
	/**
	trie := implementTrie.Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // return True
	fmt.Println(trie.Search("app"))     // return False
	fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
	*/

	// design add and search words data structure
	/**
	wordDictionary := designAddAndSearchWordsDataStructure.Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad")) // return False
	fmt.Println(wordDictionary.Search("bad")) // return True
	fmt.Println(wordDictionary.Search(".ad")) // return True
	fmt.Println(wordDictionary.Search("b..")) // return True
	*/

	// max area of island
	/**
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland.MaxAreaOfIsland(grid))

	grid1 := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland.MaxAreaOfIsland(grid1))
	*/

	// rotting orages
	// 0 representing an empty cell
	// 1 representing a fresh orange
	// 2 representing a rotten orange
	/**
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(rottingOranges.OrangesRottingBFS(grid))

	grid1 := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println(rottingOranges.OrangesRottingBFS(grid1))
	*/

	// interface test
	/**
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	*/

	// pacific altanic water flow
	/**
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlanicWaterFlow.PacificAtlantic(heights))
	*/

	// surrounded regions
	/**
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	surroundedRegions.SolveDfs(board)
	for _, row := range board {
		fmt.Printf("%s\n", row)
	}
	fmt.Printf("\n")

	board1 := [][]byte{
		{'O','X','X','O','X'},
		{'X','O','O','X','O'},
		{'X','O','X','O','X'},
		{'O','X','O','O','O'},
		{'X','X','O','X','O'},
	}
	surroundedRegions.SolveDfs(board1)
	for _, row := range board1 {
		fmt.Printf("%s\n", row)
	}
	fmt.Printf("\n")

	board2 := [][]byte {{'O'}}
	surroundedRegions.SolveDfs(board2)
	for _, row := range board2 {
		fmt.Printf("%s\n", row)
	}
	*/

	// course schedule
	/**
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(courseSchedule.CanFinishDFS(numCourses, prerequisites))

	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}, {0, 1}}
	fmt.Println(courseSchedule.CanFinishDFS(numCourses1, prerequisites1))
	*/

	// course schedule II
	/**
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(courseScheduleII.FindOrder(numCourses, prerequisites))

	numCourses1 := 4
	prerequisites1 := [][]int{{1,0}, {2,0}, {3,1},{3,2}}
	fmt.Println(courseScheduleII.FindOrder(numCourses1, prerequisites1))

	numCourses2 := 1
	prerequisites2 := [][]int{}
	fmt.Println(courseScheduleII.FindOrder(numCourses2, prerequisites2))
	*/

	// redundant connection
	/**
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(redundantConnection.FindRedundantConnectionOp(edges))
	fmt.Printf("\n")

	edges1 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	fmt.Println(redundantConnection.FindRedundantConnectionOp(edges1))
	fmt.Printf("\n")
	*/

	// min cost climbing stairs
	/**
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs.MinCostClimbingStairs(cost))

	cost1 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs.MinCostClimbingStairs(cost1))
	*/

	// palindromic substring
	/**
	s := "abc"
	fmt.Println(palindromicSubstring.CountSubstringDP(s))

	s1 := "aaa"
	fmt.Println(palindromicSubstring.CountSubstringDP(s1))
	*/

	// repeated substring pattern
	/**
	s := "aabaaba"
	fmt.Println(repeatedSubstringPattern.RepeatedSubstringPatternKMP(s))

	s1 := "bab"
	fmt.Println(repeatedSubstringPattern.RepeatedSubstringPatternKMP(s1))

	s2 := "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern.RepeatedSubstringPatternKMP (s2))
	*/

	// decode ways
	/**
	s := "12"
	fmt.Println(decodeWays.NumDecodingsTabulation(s))

	s1 := "226"
	fmt.Println(decodeWays.NumDecodingsTabulation(s1))

	s2 := "06"
	fmt.Println(decodeWays.NumDecodingsTabulation(s2))
	*/

	// ugly number
	/**
	n := 6
	fmt.Println(uglyNumber.IsUgly(n))

	n1 := 11
	fmt.Println(uglyNumber.IsUgly(n1))

	n2 := 14
	fmt.Println(uglyNumber.IsUgly(n2))
	*/

	// variable shadowing
	/**
	x := 1
	fmt.Println(x)
	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
	*/

	// maximum product subarray
	/**
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProductSubarray.MaxProduct(nums) )

	nums1 := []int{-2, 0, -1, -5}
	fmt.Println(maxProductSubarray.MaxProduct(nums1))

	nums2 := []int{0,10,10,10,10,10,10,10,10,10,-10,10,10,10,10,10,10,10,10,10,0}
	fmt.Println(maxProductSubarray.MaxProduct(nums2))
	*/

	// longest palindrome
	/**
	s := "abccccdd"
	fmt.Println(longestPalindrome.LongestPalindrome(s))

	s1 := "aa"
	fmt.Println(longestPalindrome.LongestPalindrome(s1))
	*/

	// word break
	/**
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak.WordBreak(s, wordDict))

	s1 := "applepenapple"
	wordDict1 := []string{"apple", "pen", "p", "en"}
	fmt.Println(wordBreak.WordBreak(s1, wordDict1))

	s2 := "catsandog"
	wordDict2 := []string{"cats", "dog", "sand", "and", "cat", "og"}
	fmt.Println(wordBreak.WordBreak(s2, wordDict2))
	*/

	// partition equal subset sum
	/**
	nums := []int{1, 5, 11, 5}
	fmt.Println(partitionEqualSubsetSum.CanPartitionOp(nums))

	nums1 := []int{1,2,3,5}
	fmt.Println(partitionEqualSubsetSum.CanPartitionOp(nums1))
	*/

	// min cost to connect all points
	/**
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(minCost2ConnectAllPoints.MinCostConnectPointsOp(points))
	*/

	// net work delay time
	/**
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	fmt.Println(networkDelayTime.NetworkDelayTime(times, n, k))

	times1 := [][]int{{1, 2, 1}}
	n1 := 2
	k1 := 1
	fmt.Println(networkDelayTime.NetworkDelayTime(times1, n1, k1))

	times2 := [][]int{{1, 2, 1}}
	n2 := 2
	k2 := 2
	fmt.Println(networkDelayTime.NetworkDelayTime(times2, n2, k2))
	*/

	// cheapest flights within k stops
	// flight: {from, to, price}
	// n: amount of cities
	// k: how many stops are allowed at most
	/**
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src, dst, k := 0, 3, 1
	fmt.Println(cheapestFlightsWithinKStops.FindCheapestPrice(n, flights, src, dst, k))

	n1 := 3
	flights1 := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src1, dst1, k1 := 0, 2, 1
	fmt.Println(cheapestFlightsWithinKStops.FindCheapestPrice(n1, flights1, src1, dst1, k1))

	n2 := 3
	flights2 := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src2, dst2, k2 := 0, 2, 0
	fmt.Println(cheapestFlightsWithinKStops.FindCheapestPrice(n2, flights2, src2, dst2, k2))
	*/

	// longest common sub sequence
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence.LongestCommonSubsequence(text1, text2))

	text1_1 := "abc"
	text2_1 := "def"
	fmt.Println(longestCommonSubsequence.LongestCommonSubsequence(text1_1, text2_1))
}
