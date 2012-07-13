// Copyright 2012 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package stats

import (
	"strconv"
	"testing"
	. "wolfmud.org/utils/test"
)

type result struct {
	scaled int64
	scale  string
}

var testUnsignedSubjects = []struct {
	bytes  uint64
	result result
}{
	{0, result{0, "b"}},
	{1, result{1, "b"}},
	{2, result{2, "b"}},
	{3, result{3, "b"}},
	{4, result{4, "b"}},
	{5, result{5, "b"}},
	{7, result{7, "b"}},
	{8, result{8, "b"}},
	{9, result{9, "b"}},
	{15, result{15, "b"}},
	{16, result{16, "b"}},
	{17, result{17, "b"}},
	{31, result{31, "b"}},
	{32, result{32, "b"}},
	{33, result{33, "b"}},
	{63, result{63, "b"}},
	{64, result{64, "b"}},
	{65, result{65, "b"}},
	{127, result{127, "b"}},
	{128, result{128, "b"}},
	{129, result{129, "b"}},
	{255, result{255, "b"}},
	{256, result{256, "b"}},
	{257, result{257, "b"}},
	{511, result{511, "b"}},
	{512, result{512, "b"}},
	{513, result{513, "b"}},
	{1023, result{1023, "b"}},
	{1024, result{1, "kb"}},
	{1025, result{1, "kb"}},
	{2047, result{1, "kb"}},
	{2048, result{2, "kb"}},
	{2049, result{2, "kb"}},
	{4095, result{3, "kb"}},
	{4096, result{4, "kb"}},
	{4097, result{4, "kb"}},
	{8191, result{7, "kb"}},
	{8192, result{8, "kb"}},
	{8193, result{8, "kb"}},
	{16383, result{15, "kb"}},
	{16384, result{16, "kb"}},
	{16385, result{16, "kb"}},
	{32767, result{31, "kb"}},
	{32768, result{32, "kb"}},
	{32769, result{32, "kb"}},
	{65535, result{63, "kb"}},
	{65536, result{64, "kb"}},
	{65537, result{64, "kb"}},
	{131071, result{127, "kb"}},
	{131072, result{128, "kb"}},
	{131073, result{128, "kb"}},
	{262143, result{255, "kb"}},
	{262144, result{256, "kb"}},
	{262145, result{256, "kb"}},
	{524287, result{511, "kb"}},
	{524288, result{512, "kb"}},
	{524289, result{512, "kb"}},
	{1048575, result{1023, "kb"}},
	{1048576, result{1, "Mb"}},
	{1048577, result{1, "Mb"}},
	{2097151, result{1, "Mb"}},
	{2097152, result{2, "Mb"}},
	{2097153, result{2, "Mb"}},
	{4194303, result{3, "Mb"}},
	{4194304, result{4, "Mb"}},
	{4194305, result{4, "Mb"}},
	{8388607, result{7, "Mb"}},
	{8388608, result{8, "Mb"}},
	{8388609, result{8, "Mb"}},
	{16777215, result{15, "Mb"}},
	{16777216, result{16, "Mb"}},
	{16777217, result{16, "Mb"}},
	{33554431, result{31, "Mb"}},
	{33554432, result{32, "Mb"}},
	{33554433, result{32, "Mb"}},
	{67108863, result{63, "Mb"}},
	{67108864, result{64, "Mb"}},
	{67108865, result{64, "Mb"}},
	{134217727, result{127, "Mb"}},
	{134217728, result{128, "Mb"}},
	{134217729, result{128, "Mb"}},
	{268435455, result{255, "Mb"}},
	{268435456, result{256, "Mb"}},
	{268435457, result{256, "Mb"}},
	{536870911, result{511, "Mb"}},
	{536870912, result{512, "Mb"}},
	{536870913, result{512, "Mb"}},
	{1073741823, result{1023, "Mb"}},
	{1073741824, result{1, "Gb"}},
	{1073741825, result{1, "Gb"}},
	{2147483647, result{1, "Gb"}},
	{2147483648, result{2, "Gb"}},
	{2147483649, result{2, "Gb"}},
	{4294967295, result{3, "Gb"}},
	{4294967296, result{4, "Gb"}},
	{4294967297, result{4, "Gb"}},
	{8589934591, result{7, "Gb"}},
	{8589934592, result{8, "Gb"}},
	{8589934593, result{8, "Gb"}},
	{17179869183, result{15, "Gb"}},
	{17179869184, result{16, "Gb"}},
	{17179869185, result{16, "Gb"}},
	{34359738367, result{31, "Gb"}},
	{34359738368, result{32, "Gb"}},
	{34359738369, result{32, "Gb"}},
	{68719476735, result{63, "Gb"}},
	{68719476736, result{64, "Gb"}},
	{68719476737, result{64, "Gb"}},
	{137438953471, result{127, "Gb"}},
	{137438953472, result{128, "Gb"}},
	{137438953473, result{128, "Gb"}},
	{274877906943, result{255, "Gb"}},
	{274877906944, result{256, "Gb"}},
	{274877906945, result{256, "Gb"}},
	{549755813887, result{511, "Gb"}},
	{549755813888, result{512, "Gb"}},
	{549755813889, result{512, "Gb"}},
	{1099511627775, result{1023, "Gb"}},
	{1099511627776, result{1, "Tb"}},
	{1099511627777, result{1, "Tb"}},
	{2199023255551, result{1, "Tb"}},
	{2199023255552, result{2, "Tb"}},
	{2199023255553, result{2, "Tb"}},
	{4398046511103, result{3, "Tb"}},
	{4398046511104, result{4, "Tb"}},
	{4398046511105, result{4, "Tb"}},
	{8796093022207, result{7, "Tb"}},
	{8796093022208, result{8, "Tb"}},
	{8796093022209, result{8, "Tb"}},
	{17592186044415, result{15, "Tb"}},
	{17592186044416, result{16, "Tb"}},
	{17592186044417, result{16, "Tb"}},
	{35184372088831, result{31, "Tb"}},
	{35184372088832, result{32, "Tb"}},
	{35184372088833, result{32, "Tb"}},
	{70368744177663, result{63, "Tb"}},
	{70368744177664, result{64, "Tb"}},
	{70368744177665, result{64, "Tb"}},
	{140737488355327, result{127, "Tb"}},
	{140737488355328, result{128, "Tb"}},
	{140737488355329, result{128, "Tb"}},
	{281474976710655, result{255, "Tb"}},
	{281474976710656, result{256, "Tb"}},
	{281474976710657, result{256, "Tb"}},
	{562949953421311, result{511, "Tb"}},
	{562949953421312, result{512, "Tb"}},
	{562949953421313, result{512, "Tb"}},
	{1125899906842623, result{1023, "Tb"}},
	{1125899906842624, result{1, "Pb"}},
	{1125899906842625, result{1, "Pb"}},
	{2251799813685247, result{1, "Pb"}},
	{2251799813685248, result{2, "Pb"}},
	{2251799813685249, result{2, "Pb"}},
	{4503599627370495, result{3, "Pb"}},
	{4503599627370496, result{4, "Pb"}},
	{4503599627370497, result{4, "Pb"}},
	{9007199254740991, result{7, "Pb"}},
	{9007199254740992, result{8, "Pb"}},
	{9007199254740993, result{8, "Pb"}},
	{18014398509481983, result{15, "Pb"}},
	{18014398509481984, result{16, "Pb"}},
	{18014398509481985, result{16, "Pb"}},
	{36028797018963967, result{31, "Pb"}},
	{36028797018963968, result{32, "Pb"}},
	{36028797018963969, result{32, "Pb"}},
	{72057594037927935, result{63, "Pb"}},
	{72057594037927936, result{64, "Pb"}},
	{72057594037927937, result{64, "Pb"}},
	{144115188075855871, result{127, "Pb"}},
	{144115188075855872, result{128, "Pb"}},
	{144115188075855873, result{128, "Pb"}},
	{288230376151711743, result{255, "Pb"}},
	{288230376151711744, result{256, "Pb"}},
	{288230376151711745, result{256, "Pb"}},
	{576460752303423487, result{511, "Pb"}},
	{576460752303423488, result{512, "Pb"}},
	{576460752303423489, result{512, "Pb"}},
	{1152921504606846975, result{1023, "Pb"}},
	{1152921504606846976, result{1, "Eb"}},
	{1152921504606846977, result{1, "Eb"}},
	{2305843009213693951, result{1, "Eb"}},
	{2305843009213693952, result{2, "Eb"}},
	{2305843009213693953, result{2, "Eb"}},
	{4611686018427387903, result{3, "Eb"}},
	{4611686018427387904, result{4, "Eb"}},
	{4611686018427387905, result{4, "Eb"}},
	{9223372036854775807, result{7, "Eb"}},
	{9223372036854775808, result{8, "Eb"}},
	{9223372036854775809, result{8, "Eb"}},
	{18446744073709551615, result{15, "Eb"}}, // Max value for uint64

	/* Added for completeness of the table only:

	{18446744073709551616, result{16, "Eb"}},
	{18446744073709551617, result{16, "Eb"}},

	*/
}

var testSignedSubjects = []struct {
	bytes  int64
	result result
}{
	{-9223372036854775807, result{-7, "Eb"}},
	{-4611686018427387905, result{-4, "Eb"}},
	{-4611686018427387904, result{-4, "Eb"}},
	{-4611686018427387903, result{-3, "Eb"}},
	{-2305843009213693953, result{-2, "Eb"}},
	{-2305843009213693952, result{-2, "Eb"}},
	{-2305843009213693951, result{-1, "Eb"}},
	{-1152921504606846977, result{-1, "Eb"}},
	{-1152921504606846976, result{-1, "Eb"}},
	{-1152921504606846975, result{-1023, "Pb"}},
	{-576460752303423489, result{-512, "Pb"}},
	{-576460752303423488, result{-512, "Pb"}},
	{-576460752303423487, result{-511, "Pb"}},
	{-288230376151711745, result{-256, "Pb"}},
	{-288230376151711744, result{-256, "Pb"}},
	{-288230376151711743, result{-255, "Pb"}},
	{-144115188075855873, result{-128, "Pb"}},
	{-144115188075855872, result{-128, "Pb"}},
	{-144115188075855871, result{-127, "Pb"}},
	{-72057594037927937, result{-64, "Pb"}},
	{-72057594037927936, result{-64, "Pb"}},
	{-72057594037927935, result{-63, "Pb"}},
	{-36028797018963969, result{-32, "Pb"}},
	{-36028797018963968, result{-32, "Pb"}},
	{-36028797018963967, result{-31, "Pb"}},
	{-18014398509481985, result{-16, "Pb"}},
	{-18014398509481984, result{-16, "Pb"}},
	{-18014398509481983, result{-15, "Pb"}},
	{-9007199254740993, result{-8, "Pb"}},
	{-9007199254740992, result{-8, "Pb"}},
	{-9007199254740991, result{-7, "Pb"}},
	{-4503599627370497, result{-4, "Pb"}},
	{-4503599627370496, result{-4, "Pb"}},
	{-4503599627370495, result{-3, "Pb"}},
	{-2251799813685249, result{-2, "Pb"}},
	{-2251799813685248, result{-2, "Pb"}},
	{-2251799813685247, result{-1, "Pb"}},
	{-1125899906842625, result{-1, "Pb"}},
	{-1125899906842624, result{-1, "Pb"}},
	{-1125899906842623, result{-1023, "Tb"}},
	{-562949953421313, result{-512, "Tb"}},
	{-562949953421312, result{-512, "Tb"}},
	{-562949953421311, result{-511, "Tb"}},
	{-281474976710657, result{-256, "Tb"}},
	{-281474976710656, result{-256, "Tb"}},
	{-281474976710655, result{-255, "Tb"}},
	{-140737488355329, result{-128, "Tb"}},
	{-140737488355328, result{-128, "Tb"}},
	{-140737488355327, result{-127, "Tb"}},
	{-70368744177665, result{-64, "Tb"}},
	{-70368744177664, result{-64, "Tb"}},
	{-70368744177663, result{-63, "Tb"}},
	{-35184372088833, result{-32, "Tb"}},
	{-35184372088832, result{-32, "Tb"}},
	{-35184372088831, result{-31, "Tb"}},
	{-17592186044417, result{-16, "Tb"}},
	{-17592186044416, result{-16, "Tb"}},
	{-17592186044415, result{-15, "Tb"}},
	{-8796093022209, result{-8, "Tb"}},
	{-8796093022208, result{-8, "Tb"}},
	{-8796093022207, result{-7, "Tb"}},
	{-4398046511105, result{-4, "Tb"}},
	{-4398046511104, result{-4, "Tb"}},
	{-4398046511103, result{-3, "Tb"}},
	{-2199023255553, result{-2, "Tb"}},
	{-2199023255552, result{-2, "Tb"}},
	{-2199023255551, result{-1, "Tb"}},
	{-1099511627777, result{-1, "Tb"}},
	{-1099511627776, result{-1, "Tb"}},
	{-1099511627775, result{-1023, "Gb"}},
	{-549755813889, result{-512, "Gb"}},
	{-549755813888, result{-512, "Gb"}},
	{-549755813887, result{-511, "Gb"}},
	{-274877906945, result{-256, "Gb"}},
	{-274877906944, result{-256, "Gb"}},
	{-274877906943, result{-255, "Gb"}},
	{-137438953473, result{-128, "Gb"}},
	{-137438953472, result{-128, "Gb"}},
	{-137438953471, result{-127, "Gb"}},
	{-68719476737, result{-64, "Gb"}},
	{-68719476736, result{-64, "Gb"}},
	{-68719476735, result{-63, "Gb"}},
	{-34359738369, result{-32, "Gb"}},
	{-34359738368, result{-32, "Gb"}},
	{-34359738367, result{-31, "Gb"}},
	{-17179869185, result{-16, "Gb"}},
	{-17179869184, result{-16, "Gb"}},
	{-17179869183, result{-15, "Gb"}},
	{-8589934593, result{-8, "Gb"}},
	{-8589934592, result{-8, "Gb"}},
	{-8589934591, result{-7, "Gb"}},
	{-4294967297, result{-4, "Gb"}},
	{-4294967296, result{-4, "Gb"}},
	{-4294967295, result{-3, "Gb"}},
	{-2147483649, result{-2, "Gb"}},
	{-2147483648, result{-2, "Gb"}},
	{-2147483647, result{-1, "Gb"}},
	{-1073741825, result{-1, "Gb"}},
	{-1073741824, result{-1, "Gb"}},
	{-1073741823, result{-1023, "Mb"}},
	{-536870913, result{-512, "Mb"}},
	{-536870912, result{-512, "Mb"}},
	{-536870911, result{-511, "Mb"}},
	{-268435457, result{-256, "Mb"}},
	{-268435456, result{-256, "Mb"}},
	{-268435455, result{-255, "Mb"}},
	{-134217729, result{-128, "Mb"}},
	{-134217728, result{-128, "Mb"}},
	{-134217727, result{-127, "Mb"}},
	{-67108865, result{-64, "Mb"}},
	{-67108864, result{-64, "Mb"}},
	{-67108863, result{-63, "Mb"}},
	{-33554433, result{-32, "Mb"}},
	{-33554432, result{-32, "Mb"}},
	{-33554431, result{-31, "Mb"}},
	{-16777217, result{-16, "Mb"}},
	{-16777216, result{-16, "Mb"}},
	{-16777215, result{-15, "Mb"}},
	{-8388609, result{-8, "Mb"}},
	{-8388608, result{-8, "Mb"}},
	{-8388607, result{-7, "Mb"}},
	{-4194305, result{-4, "Mb"}},
	{-4194304, result{-4, "Mb"}},
	{-4194303, result{-3, "Mb"}},
	{-2097153, result{-2, "Mb"}},
	{-2097152, result{-2, "Mb"}},
	{-2097151, result{-1, "Mb"}},
	{-1048577, result{-1, "Mb"}},
	{-1048576, result{-1, "Mb"}},
	{-1048575, result{-1023, "kb"}},
	{-524289, result{-512, "kb"}},
	{-524288, result{-512, "kb"}},
	{-524287, result{-511, "kb"}},
	{-262145, result{-256, "kb"}},
	{-262144, result{-256, "kb"}},
	{-262143, result{-255, "kb"}},
	{-131073, result{-128, "kb"}},
	{-131072, result{-128, "kb"}},
	{-131071, result{-127, "kb"}},
	{-65537, result{-64, "kb"}},
	{-65536, result{-64, "kb"}},
	{-65535, result{-63, "kb"}},
	{-32769, result{-32, "kb"}},
	{-32768, result{-32, "kb"}},
	{-32767, result{-31, "kb"}},
	{-16385, result{-16, "kb"}},
	{-16384, result{-16, "kb"}},
	{-16383, result{-15, "kb"}},
	{-8193, result{-8, "kb"}},
	{-8192, result{-8, "kb"}},
	{-8191, result{-7, "kb"}},
	{-4097, result{-4, "kb"}},
	{-4096, result{-4, "kb"}},
	{-4095, result{-3, "kb"}},
	{-2049, result{-2, "kb"}},
	{-2048, result{-2, "kb"}},
	{-2047, result{-1, "kb"}},
	{-1025, result{-1, "kb"}},
	{-1024, result{-1, "kb"}},
	{-1023, result{-1023, "b"}},
	{-513, result{-513, "b"}},
	{-512, result{-512, "b"}},
	{-511, result{-511, "b"}},
	{-257, result{-257, "b"}},
	{-256, result{-256, "b"}},
	{-255, result{-255, "b"}},
	{-129, result{-129, "b"}},
	{-128, result{-128, "b"}},
	{-127, result{-127, "b"}},
	{-65, result{-65, "b"}},
	{-64, result{-64, "b"}},
	{-63, result{-63, "b"}},
	{-33, result{-33, "b"}},
	{-32, result{-32, "b"}},
	{-31, result{-31, "b"}},
	{-17, result{-17, "b"}},
	{-16, result{-16, "b"}},
	{-15, result{-15, "b"}},
	{-9, result{-9, "b"}},
	{-8, result{-8, "b"}},
	{-7, result{-7, "b"}},
	{-5, result{-5, "b"}},
	{-4, result{-4, "b"}},
	{-3, result{-3, "b"}},
	{-2, result{-2, "b"}},
	{-1, result{-1, "b"}},
	{-0, result{-0, "b"}},
	{0, result{0, "b"}},
	{1, result{1, "b"}},
	{2, result{2, "b"}},
	{3, result{3, "b"}},
	{4, result{4, "b"}},
	{5, result{5, "b"}},
	{7, result{7, "b"}},
	{8, result{8, "b"}},
	{9, result{9, "b"}},
	{15, result{15, "b"}},
	{16, result{16, "b"}},
	{17, result{17, "b"}},
	{31, result{31, "b"}},
	{32, result{32, "b"}},
	{33, result{33, "b"}},
	{63, result{63, "b"}},
	{64, result{64, "b"}},
	{65, result{65, "b"}},
	{127, result{127, "b"}},
	{128, result{128, "b"}},
	{129, result{129, "b"}},
	{255, result{255, "b"}},
	{256, result{256, "b"}},
	{257, result{257, "b"}},
	{511, result{511, "b"}},
	{512, result{512, "b"}},
	{513, result{513, "b"}},
	{1023, result{1023, "b"}},
	{1024, result{1, "kb"}},
	{1025, result{1, "kb"}},
	{2047, result{1, "kb"}},
	{2048, result{2, "kb"}},
	{2049, result{2, "kb"}},
	{4095, result{3, "kb"}},
	{4096, result{4, "kb"}},
	{4097, result{4, "kb"}},
	{8191, result{7, "kb"}},
	{8192, result{8, "kb"}},
	{8193, result{8, "kb"}},
	{16383, result{15, "kb"}},
	{16384, result{16, "kb"}},
	{16385, result{16, "kb"}},
	{32767, result{31, "kb"}},
	{32768, result{32, "kb"}},
	{32769, result{32, "kb"}},
	{65535, result{63, "kb"}},
	{65536, result{64, "kb"}},
	{65537, result{64, "kb"}},
	{131071, result{127, "kb"}},
	{131072, result{128, "kb"}},
	{131073, result{128, "kb"}},
	{262143, result{255, "kb"}},
	{262144, result{256, "kb"}},
	{262145, result{256, "kb"}},
	{524287, result{511, "kb"}},
	{524288, result{512, "kb"}},
	{524289, result{512, "kb"}},
	{1048575, result{1023, "kb"}},
	{1048576, result{1, "Mb"}},
	{1048577, result{1, "Mb"}},
	{2097151, result{1, "Mb"}},
	{2097152, result{2, "Mb"}},
	{2097153, result{2, "Mb"}},
	{4194303, result{3, "Mb"}},
	{4194304, result{4, "Mb"}},
	{4194305, result{4, "Mb"}},
	{8388607, result{7, "Mb"}},
	{8388608, result{8, "Mb"}},
	{8388609, result{8, "Mb"}},
	{16777215, result{15, "Mb"}},
	{16777216, result{16, "Mb"}},
	{16777217, result{16, "Mb"}},
	{33554431, result{31, "Mb"}},
	{33554432, result{32, "Mb"}},
	{33554433, result{32, "Mb"}},
	{67108863, result{63, "Mb"}},
	{67108864, result{64, "Mb"}},
	{67108865, result{64, "Mb"}},
	{134217727, result{127, "Mb"}},
	{134217728, result{128, "Mb"}},
	{134217729, result{128, "Mb"}},
	{268435455, result{255, "Mb"}},
	{268435456, result{256, "Mb"}},
	{268435457, result{256, "Mb"}},
	{536870911, result{511, "Mb"}},
	{536870912, result{512, "Mb"}},
	{536870913, result{512, "Mb"}},
	{1073741823, result{1023, "Mb"}},
	{1073741824, result{1, "Gb"}},
	{1073741825, result{1, "Gb"}},
	{2147483647, result{1, "Gb"}},
	{2147483648, result{2, "Gb"}},
	{2147483649, result{2, "Gb"}},
	{4294967295, result{3, "Gb"}},
	{4294967296, result{4, "Gb"}},
	{4294967297, result{4, "Gb"}},
	{8589934591, result{7, "Gb"}},
	{8589934592, result{8, "Gb"}},
	{8589934593, result{8, "Gb"}},
	{17179869183, result{15, "Gb"}},
	{17179869184, result{16, "Gb"}},
	{17179869185, result{16, "Gb"}},
	{34359738367, result{31, "Gb"}},
	{34359738368, result{32, "Gb"}},
	{34359738369, result{32, "Gb"}},
	{68719476735, result{63, "Gb"}},
	{68719476736, result{64, "Gb"}},
	{68719476737, result{64, "Gb"}},
	{137438953471, result{127, "Gb"}},
	{137438953472, result{128, "Gb"}},
	{137438953473, result{128, "Gb"}},
	{274877906943, result{255, "Gb"}},
	{274877906944, result{256, "Gb"}},
	{274877906945, result{256, "Gb"}},
	{549755813887, result{511, "Gb"}},
	{549755813888, result{512, "Gb"}},
	{549755813889, result{512, "Gb"}},
	{1099511627775, result{1023, "Gb"}},
	{1099511627776, result{1, "Tb"}},
	{1099511627777, result{1, "Tb"}},
	{2199023255551, result{1, "Tb"}},
	{2199023255552, result{2, "Tb"}},
	{2199023255553, result{2, "Tb"}},
	{4398046511103, result{3, "Tb"}},
	{4398046511104, result{4, "Tb"}},
	{4398046511105, result{4, "Tb"}},
	{8796093022207, result{7, "Tb"}},
	{8796093022208, result{8, "Tb"}},
	{8796093022209, result{8, "Tb"}},
	{17592186044415, result{15, "Tb"}},
	{17592186044416, result{16, "Tb"}},
	{17592186044417, result{16, "Tb"}},
	{35184372088831, result{31, "Tb"}},
	{35184372088832, result{32, "Tb"}},
	{35184372088833, result{32, "Tb"}},
	{70368744177663, result{63, "Tb"}},
	{70368744177664, result{64, "Tb"}},
	{70368744177665, result{64, "Tb"}},
	{140737488355327, result{127, "Tb"}},
	{140737488355328, result{128, "Tb"}},
	{140737488355329, result{128, "Tb"}},
	{281474976710655, result{255, "Tb"}},
	{281474976710656, result{256, "Tb"}},
	{281474976710657, result{256, "Tb"}},
	{562949953421311, result{511, "Tb"}},
	{562949953421312, result{512, "Tb"}},
	{562949953421313, result{512, "Tb"}},
	{1125899906842623, result{1023, "Tb"}},
	{1125899906842624, result{1, "Pb"}},
	{1125899906842625, result{1, "Pb"}},
	{2251799813685247, result{1, "Pb"}},
	{2251799813685248, result{2, "Pb"}},
	{2251799813685249, result{2, "Pb"}},
	{4503599627370495, result{3, "Pb"}},
	{4503599627370496, result{4, "Pb"}},
	{4503599627370497, result{4, "Pb"}},
	{9007199254740991, result{7, "Pb"}},
	{9007199254740992, result{8, "Pb"}},
	{9007199254740993, result{8, "Pb"}},
	{18014398509481983, result{15, "Pb"}},
	{18014398509481984, result{16, "Pb"}},
	{18014398509481985, result{16, "Pb"}},
	{36028797018963967, result{31, "Pb"}},
	{36028797018963968, result{32, "Pb"}},
	{36028797018963969, result{32, "Pb"}},
	{72057594037927935, result{63, "Pb"}},
	{72057594037927936, result{64, "Pb"}},
	{72057594037927937, result{64, "Pb"}},
	{144115188075855871, result{127, "Pb"}},
	{144115188075855872, result{128, "Pb"}},
	{144115188075855873, result{128, "Pb"}},
	{288230376151711743, result{255, "Pb"}},
	{288230376151711744, result{256, "Pb"}},
	{288230376151711745, result{256, "Pb"}},
	{576460752303423487, result{511, "Pb"}},
	{576460752303423488, result{512, "Pb"}},
	{576460752303423489, result{512, "Pb"}},
	{1152921504606846975, result{1023, "Pb"}},
	{1152921504606846976, result{1, "Eb"}},
	{1152921504606846977, result{1, "Eb"}},
	{2305843009213693951, result{1, "Eb"}},
	{2305843009213693952, result{2, "Eb"}},
	{2305843009213693953, result{2, "Eb"}},
	{4611686018427387903, result{3, "Eb"}},
	{4611686018427387904, result{4, "Eb"}},
	{4611686018427387905, result{4, "Eb"}},
	{9223372036854775807, result{7, "Eb"}},
}

func TestUscale(t *testing.T) {
	for _, s := range testUnsignedSubjects {
		r := result{}
		r.scaled, r.scale = uscale(s.bytes)
		Equal(t, "uscale with "+strconv.FormatUint(s.bytes, 10)+" bytes", s.result, r)
	}
}

func TestScale(t *testing.T) {
	for _, s := range testSignedSubjects {
		r := result{}
		r.scaled, r.scale = scale(s.bytes)
		Equal(t, "scale with "+strconv.FormatInt(s.bytes, 10)+" bytes", s.result, r)
	}
}
