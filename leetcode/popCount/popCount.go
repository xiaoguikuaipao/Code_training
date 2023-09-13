package popCount

//表的预计算，256表示8个bit表示的数范围
var pc [256]byte

//256个元素代表从00000000到11111111每个值，pc[i]表示一个i [0,255]所代表的二进制数中含1的个数
//pc[0]=0, pc[1]=1。然后通过动态规划构建整个表

func init() {
	for i := range pc {
		//5(0101)的含1数是2(0010)的含1数+位数
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 不断截取低8位，找对应字节的含1数，最后加起来8个字节的含1数
func PopCount(x uint64) int {
	//byte()会直接截断低8位, 这个数是[0, 255]，直接查表知道这个数对应的含1数,就不用for循环一个个数
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
