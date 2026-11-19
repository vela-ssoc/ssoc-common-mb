package model

import (
	"cmp"
	"encoding/binary"
	"strconv"
	"strings"
)

// SemVer https://semver.org/lang/zh-CN/
type SemVer struct {
	Literal string `bson:"literal" json:"literal"` // 原始版本字面值，如 "2.3.2"
	Weight  uint64 `bson:"weight"  json:"weight"`  // 压缩后的排序权重，仅用于比较
}

func (s SemVer) Compare(b SemVer) int {
	return cmp.Compare(s.Weight, b.Weight)
}

func (s SemVer) String() string {
	return s.Literal
}

// ParseSemVer 解析版本号。
//
// FIXME 这种方式仅适用于特定环境下的版本号，不具备普适性，需要调用者提前校验好 literal。
//
//	版本号 X.Y.Z 中 X Y Z 均要大于等于 0 小于等于 255 才可用。
//	prerelease 至多保留 5 个字符。
//
// 内存布局： MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]
//
// 0       1       2       3       4       5       6       7       8
// +-------+-------+-------+-------+-------+-------+-------+-------+
// | MAJOR | MINOR | PATCH |          PRERELEASE (5 bytes)         |
// +-------+-------+-------+-------+-------+-------+-------+-------+
func ParseSemVer(literal string) SemVer {
	sv := SemVer{Literal: literal}

	// 版本号 MAJOR.MINOR.PATCH 位按照点分十进制依次比较。
	//
	// - 号后面是 PRERELEASE，PRERELEASE 按照字典序大小比较，
	// 所有带 PRERELEASE 的版本号，都要小于主版本号，例如：
	// 0.0.1 > 0.0.1-alpha
	// 0.0.1 > 0.0.1-balabala
	//
	// + 号后面的是 BUILD，不影响版本优先级比较。例如：
	// 0.0.1 == 0.0.1+balabala 大小是一样的。
	literal, _, _ = strings.Cut(literal, "+") // 移除 BUILD 部分
	semver, prerelease, _ := strings.Cut(literal, "-")
	bucket := make([]byte, 8) // uint64
	for i, v := range strings.SplitN(semver, ".", 3) {
		n, _ := strconv.Atoi(v) // 假定是符合预期的
		bucket[i] = byte(n)
	}

	idx := 3 // PRERELEASE 剩余位置
	if prerelease == "" {
		// 无 PRERELEASE 说明是主版本号，填充 0xff 保证结果最大。
		for ; idx < len(bucket); idx++ {
			bucket[idx] = 0xff
		}
		sv.Weight = binary.BigEndian.Uint64(bucket)

		return sv
	}

	for _, b := range []byte(prerelease) {
		if idx >= len(bucket) { // 满了
			break
		}
		bucket[idx] = b
		idx++
	}
	sv.Weight = binary.BigEndian.Uint64(bucket)

	return sv
}
