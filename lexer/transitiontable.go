
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 33 : // ['!','!']
				return 2
			case r == 38 : // ['&','&']
				return 3
			case r == 40 : // ['(','(']
				return 4
			case r == 41 : // [')',')']
				return 5
			case r == 42 : // ['*','*']
				return 6
			case r == 43 : // ['+','+']
				return 7
			case r == 44 : // [',',',']
				return 8
			case r == 45 : // ['-','-']
				return 9
			case r == 46 : // ['.','.']
				return 10
			case r == 47 : // ['/','/']
				return 11
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case r == 58 : // [':',':']
				return 13
			case r == 59 : // [';',';']
				return 14
			case r == 60 : // ['<','<']
				return 15
			case r == 61 : // ['=','=']
				return 16
			case r == 62 : // ['>','>']
				return 17
			case r == 91 : // ['[','[']
				return 18
			case r == 93 : // [']',']']
				return 19
			case 97 <= r && r <= 100 : // ['a','d']
				return 20
			case r == 101 : // ['e','e']
				return 21
			case r == 102 : // ['f','f']
				return 22
			case 103 <= r && r <= 104 : // ['g','h']
				return 20
			case r == 105 : // ['i','i']
				return 23
			case 106 <= r && r <= 113 : // ['j','q']
				return 20
			case r == 114 : // ['r','r']
				return 24
			case r == 115 : // ['s','s']
				return 20
			case r == 116 : // ['t','t']
				return 25
			case r == 117 : // ['u','u']
				return 20
			case r == 118 : // ['v','v']
				return 26
			case r == 119 : // ['w','w']
				return 27
			case 120 <= r && r <= 122 : // ['x','z']
				return 20
			case r == 123 : // ['{','{']
				return 28
			case r == 124 : // ['|','|']
				return 29
			case r == 125 : // ['}','}']
				return 30
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 31
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 38 : // ['&','&']
				return 32
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 43 : // ['+','+']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 34
			case r == 62 : // ['>','>']
				return 35
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 36
			case r == 47 : // ['/','/']
				return 37
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 38
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 107 : // ['a','k']
				return 39
			case r == 108 : // ['l','l']
				return 40
			case 109 <= r && r <= 122 : // ['m','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case 49 <= r && r <= 57 : // ['1','9']
				return 41
			case r == 95 : // ['_','_']
				return 39
			case r == 97 : // ['a','a']
				return 42
			case 98 <= r && r <= 110 : // ['b','n']
				return 39
			case r == 111 : // ['o','o']
				return 43
			case 112 <= r && r <= 116 : // ['p','t']
				return 39
			case r == 117 : // ['u','u']
				return 44
			case 118 <= r && r <= 122 : // ['v','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 101 : // ['a','e']
				return 39
			case r == 102 : // ['f','f']
				return 45
			case 103 <= r && r <= 109 : // ['g','m']
				return 39
			case r == 110 : // ['n','n']
				return 46
			case 111 <= r && r <= 122 : // ['o','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 47
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 39
			case r == 114 : // ['r','r']
				return 48
			case 115 <= r && r <= 122 : // ['s','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case r == 97 : // ['a','a']
				return 49
			case 98 <= r && r <= 122 : // ['b','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 103 : // ['a','g']
				return 39
			case r == 104 : // ['h','h']
				return 50
			case 105 <= r && r <= 122 : // ['i','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 124 : // ['|','|']
				return 51
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 52
			
			
			default:
				return 36
			}
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 53
			case r == 13 : // ['\r','\r']
				return 53
			
			
			default:
				return 37
			}
			
		},
	
		// S38
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 114 : // ['a','r']
				return 39
			case r == 115 : // ['s','s']
				return 54
			case 116 <= r && r <= 122 : // ['t','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case 48 <= r && r <= 54 : // ['0','6']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 107 : // ['a','k']
				return 39
			case r == 108 : // ['l','l']
				return 56
			case 109 <= r && r <= 122 : // ['m','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 39
			case r == 114 : // ['r','r']
				return 57
			case 115 <= r && r <= 122 : // ['s','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 109 : // ['a','m']
				return 39
			case r == 110 : // ['n','n']
				return 58
			case 111 <= r && r <= 122 : // ['o','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 111 : // ['a','o']
				return 39
			case r == 112 : // ['p','p']
				return 59
			case 113 <= r && r <= 122 : // ['q','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 39
			case r == 116 : // ['t','t']
				return 60
			case 117 <= r && r <= 122 : // ['u','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 116 : // ['a','t']
				return 39
			case r == 117 : // ['u','u']
				return 61
			case 118 <= r && r <= 122 : // ['v','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 39
			case r == 114 : // ['r','r']
				return 62
			case 115 <= r && r <= 122 : // ['s','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 104 : // ['a','h']
				return 39
			case r == 105 : // ['i','i']
				return 63
			case 106 <= r && r <= 122 : // ['j','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 52
			case r == 47 : // ['/','/']
				return 64
			
			
			default:
				return 36
			}
			
		},
	
		// S53
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 65
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 114 : // ['a','r']
				return 39
			case r == 115 : // ['s','s']
				return 66
			case 116 <= r && r <= 122 : // ['t','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 67
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 98 : // ['a','b']
				return 39
			case r == 99 : // ['c','c']
				return 68
			case 100 <= r && r <= 122 : // ['d','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 116 : // ['a','t']
				return 39
			case r == 117 : // ['u','u']
				return 69
			case 118 <= r && r <= 122 : // ['v','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 116 : // ['a','t']
				return 39
			case r == 117 : // ['u','u']
				return 70
			case 118 <= r && r <= 122 : // ['v','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 71
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 72
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 107 : // ['a','k']
				return 39
			case r == 108 : // ['l','l']
				return 73
			case 109 <= r && r <= 122 : // ['m','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 74
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case r == 97 : // ['a','a']
				return 75
			case 98 <= r && r <= 122 : // ['b','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 39
			case r == 116 : // ['t','t']
				return 76
			case 117 <= r && r <= 122 : // ['u','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S69
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 39
			case r == 116 : // ['t','t']
				return 77
			case 117 <= r && r <= 122 : // ['u','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S70
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 39
			case r == 114 : // ['r','r']
				return 78
			case 115 <= r && r <= 122 : // ['s','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S71
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S72
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 111 : // ['a','o']
				return 79
			case 112 <= r && r <= 122 : // ['p','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S73
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 39
			case r == 101 : // ['e','e']
				return 80
			case 102 <= r && r <= 122 : // ['f','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S74
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S75
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 98 : // ['a','b']
				return 39
			case r == 99 : // ['c','c']
				return 81
			case 100 <= r && r <= 122 : // ['d','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S76
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 104 : // ['a','h']
				return 39
			case r == 105 : // ['i','i']
				return 82
			case 106 <= r && r <= 122 : // ['j','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S77
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S78
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 109 : // ['a','m']
				return 39
			case r == 110 : // ['n','n']
				return 83
			case 111 <= r && r <= 122 : // ['o','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S79
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S80
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S81
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 103 : // ['a','g']
				return 39
			case r == 104 : // ['h','h']
				return 84
			case 105 <= r && r <= 122 : // ['i','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S82
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 110 : // ['a','n']
				return 39
			case r == 111 : // ['o','o']
				return 85
			case 112 <= r && r <= 122 : // ['p','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S83
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S84
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S85
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 109 : // ['a','m']
				return 39
			case r == 110 : // ['n','n']
				return 86
			case 111 <= r && r <= 122 : // ['o','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S86
		func(r rune) int {
			switch {
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
			
		},
	
}
