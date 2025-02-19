// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/johnkerl/miller/internal/pkg/parsing/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 100,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 95,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 85,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 25,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 26,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 92,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 88,
		Ignore: "",
	},
	ActionRow{ // S12
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S13
		Accept: 89,
		Ignore: "",
	},
	ActionRow{ // S14
		Accept: 99,
		Ignore: "",
	},
	ActionRow{ // S15
		Accept: 93,
		Ignore: "",
	},
	ActionRow{ // S16
		Accept: 105,
		Ignore: "",
	},
	ActionRow{ // S17
		Accept: 105,
		Ignore: "",
	},
	ActionRow{ // S18
		Accept: 70,
		Ignore: "",
	},
	ActionRow{ // S19
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S20
		Accept: 82,
		Ignore: "",
	},
	ActionRow{ // S21
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S22
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S23
		Accept: 69,
		Ignore: "",
	},
	ActionRow{ // S24
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S25
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S26
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S27
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S28
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S29
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S30
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S31
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S32
		Accept: 114,
		Ignore: "",
	},
	ActionRow{ // S33
		Accept: 30,
		Ignore: "",
	},
	ActionRow{ // S34
		Accept: 84,
		Ignore: "",
	},
	ActionRow{ // S35
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S36
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S37
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S38
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S39
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S40
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S41
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S42
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S43
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S44
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S45
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S46
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S47
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S48
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S49
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S50
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S51
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S52
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S53
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S54
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S55
		Accept: 101,
		Ignore: "",
	},
	ActionRow{ // S56
		Accept: 79,
		Ignore: "",
	},
	ActionRow{ // S57
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S58
		Accept: 103,
		Ignore: "",
	},
	ActionRow{ // S59
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S60
		Accept: -1,
		Ignore: "!comment",
	},
	ActionRow{ // S61
		Accept: 34,
		Ignore: "",
	},
	ActionRow{ // S62
		Accept: 28,
		Ignore: "",
	},
	ActionRow{ // S63
		Accept: 28,
		Ignore: "",
	},
	ActionRow{ // S64
		Accept: 29,
		Ignore: "",
	},
	ActionRow{ // S65
		Accept: 28,
		Ignore: "",
	},
	ActionRow{ // S66
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S67
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S68
		Accept: 67,
		Ignore: "",
	},
	ActionRow{ // S69
		Accept: 73,
		Ignore: "",
	},
	ActionRow{ // S70
		Accept: 56,
		Ignore: "",
	},
	ActionRow{ // S71
		Accept: 102,
		Ignore: "",
	},
	ActionRow{ // S72
		Accept: 64,
		Ignore: "",
	},
	ActionRow{ // S73
		Accept: 61,
		Ignore: "",
	},
	ActionRow{ // S74
		Accept: 63,
		Ignore: "",
	},
	ActionRow{ // S75
		Accept: 96,
		Ignore: "",
	},
	ActionRow{ // S76
		Accept: 90,
		Ignore: "",
	},
	ActionRow{ // S77
		Accept: 91,
		Ignore: "",
	},
	ActionRow{ // S78
		Accept: 97,
		Ignore: "",
	},
	ActionRow{ // S79
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S80
		Accept: 62,
		Ignore: "",
	},
	ActionRow{ // S81
		Accept: 94,
		Ignore: "",
	},
	ActionRow{ // S82
		Accept: 65,
		Ignore: "",
	},
	ActionRow{ // S83
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S84
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S85
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S86
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S87
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S88
		Accept: 86,
		Ignore: "",
	},
	ActionRow{ // S89
		Accept: 83,
		Ignore: "",
	},
	ActionRow{ // S90
		Accept: 78,
		Ignore: "",
	},
	ActionRow{ // S91
		Accept: 76,
		Ignore: "",
	},
	ActionRow{ // S92
		Accept: 81,
		Ignore: "",
	},
	ActionRow{ // S93
		Accept: 10,
		Ignore: "",
	},
	ActionRow{ // S94
		Accept: 74,
		Ignore: "",
	},
	ActionRow{ // S95
		Accept: 38,
		Ignore: "",
	},
	ActionRow{ // S96
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S97
		Accept: 36,
		Ignore: "",
	},
	ActionRow{ // S98
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S99
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S100
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S101
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S102
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S103
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S104
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S105
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S106
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S107
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S108
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S109
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S110
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S111
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S112
		Accept: 122,
		Ignore: "",
	},
	ActionRow{ // S113
		Accept: 123,
		Ignore: "",
	},
	ActionRow{ // S114
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S115
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S116
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S117
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S118
		Accept: 128,
		Ignore: "",
	},
	ActionRow{ // S119
		Accept: 57,
		Ignore: "",
	},
	ActionRow{ // S120
		Accept: 72,
		Ignore: "",
	},
	ActionRow{ // S121
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S122
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S123
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S124
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S125
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S126
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S127
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S128
		Accept: 137,
		Ignore: "",
	},
	ActionRow{ // S129
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S130
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S131
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S132
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S133
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S134
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S135
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S136
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S137
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S138
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S139
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S140
		Accept: 133,
		Ignore: "",
	},
	ActionRow{ // S141
		Accept: 139,
		Ignore: "",
	},
	ActionRow{ // S142
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S143
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S144
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S145
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S146
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S147
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S148
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S149
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S150
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S151
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S152
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S153
		Accept: 55,
		Ignore: "",
	},
	ActionRow{ // S154
		Accept: 71,
		Ignore: "",
	},
	ActionRow{ // S155
		Accept: 77,
		Ignore: "",
	},
	ActionRow{ // S156
		Accept: 104,
		Ignore: "",
	},
	ActionRow{ // S157
		Accept: 103,
		Ignore: "",
	},
	ActionRow{ // S158
		Accept: 32,
		Ignore: "",
	},
	ActionRow{ // S159
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S160
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S161
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S162
		Accept: 52,
		Ignore: "",
	},
	ActionRow{ // S163
		Accept: 68,
		Ignore: "",
	},
	ActionRow{ // S164
		Accept: 98,
		Ignore: "",
	},
	ActionRow{ // S165
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S166
		Accept: 66,
		Ignore: "",
	},
	ActionRow{ // S167
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S168
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S169
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S170
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S171
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S172
		Accept: 105,
		Ignore: "",
	},
	ActionRow{ // S173
		Accept: 105,
		Ignore: "",
	},
	ActionRow{ // S174
		Accept: 105,
		Ignore: "",
	},
	ActionRow{ // S175
		Accept: 58,
		Ignore: "",
	},
	ActionRow{ // S176
		Accept: 80,
		Ignore: "",
	},
	ActionRow{ // S177
		Accept: 59,
		Ignore: "",
	},
	ActionRow{ // S178
		Accept: 87,
		Ignore: "",
	},
	ActionRow{ // S179
		Accept: 53,
		Ignore: "",
	},
	ActionRow{ // S180
		Accept: 75,
		Ignore: "",
	},
	ActionRow{ // S181
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S182
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S183
		Accept: 35,
		Ignore: "",
	},
	ActionRow{ // S184
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S185
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S186
		Accept: 127,
		Ignore: "",
	},
	ActionRow{ // S187
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S188
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S189
		Accept: 124,
		Ignore: "",
	},
	ActionRow{ // S190
		Accept: 116,
		Ignore: "",
	},
	ActionRow{ // S191
		Accept: 115,
		Ignore: "",
	},
	ActionRow{ // S192
		Accept: 117,
		Ignore: "",
	},
	ActionRow{ // S193
		Accept: 109,
		Ignore: "",
	},
	ActionRow{ // S194
		Accept: 112,
		Ignore: "",
	},
	ActionRow{ // S195
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S196
		Accept: 110,
		Ignore: "",
	},
	ActionRow{ // S197
		Accept: 119,
		Ignore: "",
	},
	ActionRow{ // S198
		Accept: 118,
		Ignore: "",
	},
	ActionRow{ // S199
		Accept: 120,
		Ignore: "",
	},
	ActionRow{ // S200
		Accept: 129,
		Ignore: "",
	},
	ActionRow{ // S201
		Accept: 51,
		Ignore: "",
	},
	ActionRow{ // S202
		Accept: 39,
		Ignore: "",
	},
	ActionRow{ // S203
		Accept: 41,
		Ignore: "",
	},
	ActionRow{ // S204
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S205
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S206
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S207
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S208
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S209
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S210
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S211
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S212
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S213
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S214
		Accept: 132,
		Ignore: "",
	},
	ActionRow{ // S215
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S216
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S217
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S218
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S219
		Accept: 138,
		Ignore: "",
	},
	ActionRow{ // S220
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S221
		Accept: 44,
		Ignore: "",
	},
	ActionRow{ // S222
		Accept: 45,
		Ignore: "",
	},
	ActionRow{ // S223
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S224
		Accept: 46,
		Ignore: "",
	},
	ActionRow{ // S225
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S226
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S227
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S228
		Accept: 47,
		Ignore: "",
	},
	ActionRow{ // S229
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S230
		Accept: 21,
		Ignore: "",
	},
	ActionRow{ // S231
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S232
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S233
		Accept: 48,
		Ignore: "",
	},
	ActionRow{ // S234
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S235
		Accept: 50,
		Ignore: "",
	},
	ActionRow{ // S236
		Accept: 104,
		Ignore: "",
	},
	ActionRow{ // S237
		Accept: 33,
		Ignore: "",
	},
	ActionRow{ // S238
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S239
		Accept: 31,
		Ignore: "",
	},
	ActionRow{ // S240
		Accept: 31,
		Ignore: "",
	},
	ActionRow{ // S241
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S242
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S243
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S244
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S245
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S246
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S247
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S248
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S249
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S250
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S251
		Accept: 60,
		Ignore: "",
	},
	ActionRow{ // S252
		Accept: 54,
		Ignore: "",
	},
	ActionRow{ // S253
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S254
		Accept: 37,
		Ignore: "",
	},
	ActionRow{ // S255
		Accept: 37,
		Ignore: "",
	},
	ActionRow{ // S256
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S257
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S258
		Accept: 111,
		Ignore: "",
	},
	ActionRow{ // S259
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S260
		Accept: 42,
		Ignore: "",
	},
	ActionRow{ // S261
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S262
		Accept: 130,
		Ignore: "",
	},
	ActionRow{ // S263
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S264
		Accept: 19,
		Ignore: "",
	},
	ActionRow{ // S265
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S266
		Accept: 134,
		Ignore: "",
	},
	ActionRow{ // S267
		Accept: 135,
		Ignore: "",
	},
	ActionRow{ // S268
		Accept: 24,
		Ignore: "",
	},
	ActionRow{ // S269
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S270
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S271
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S272
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S273
		Accept: 142,
		Ignore: "",
	},
	ActionRow{ // S274
		Accept: 108,
		Ignore: "",
	},
	ActionRow{ // S275
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S276
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S277
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S278
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S279
		Accept: 143,
		Ignore: "",
	},
	ActionRow{ // S280
		Accept: 107,
		Ignore: "",
	},
	ActionRow{ // S281
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S282
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S283
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S284
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S285
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S286
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S287
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S288
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S289
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S290
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S291
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S292
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S293
		Accept: 131,
		Ignore: "",
	},
	ActionRow{ // S294
		Accept: 140,
		Ignore: "",
	},
	ActionRow{ // S295
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S296
		Accept: 20,
		Ignore: "",
	},
	ActionRow{ // S297
		Accept: 23,
		Ignore: "",
	},
	ActionRow{ // S298
		Accept: 22,
		Ignore: "",
	},
	ActionRow{ // S299
		Accept: 27,
		Ignore: "",
	},
	ActionRow{ // S300
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S301
		Accept: 107,
		Ignore: "",
	},
	ActionRow{ // S302
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S303
		Accept: 43,
		Ignore: "",
	},
	ActionRow{ // S304
		Accept: 49,
		Ignore: "",
	},
	ActionRow{ // S305
		Accept: 14,
		Ignore: "",
	},
	ActionRow{ // S306
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S307
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S308
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S309
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S310
		Accept: 136,
		Ignore: "",
	},
	ActionRow{ // S311
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S312
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S313
		Accept: 106,
		Ignore: "",
	},
	ActionRow{ // S314
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S315
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S316
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S317
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S318
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S319
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S320
		Accept: 16,
		Ignore: "",
	},
	ActionRow{ // S321
		Accept: 144,
		Ignore: "",
	},
	ActionRow{ // S322
		Accept: 13,
		Ignore: "",
	},
	ActionRow{ // S323
		Accept: 12,
		Ignore: "",
	},
	ActionRow{ // S324
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S325
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S326
		Accept: 126,
		Ignore: "",
	},
	ActionRow{ // S327
		Accept: 121,
		Ignore: "",
	},
	ActionRow{ // S328
		Accept: 40,
		Ignore: "",
	},
	ActionRow{ // S329
		Accept: 18,
		Ignore: "",
	},
	ActionRow{ // S330
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S331
		Accept: 125,
		Ignore: "",
	},
	ActionRow{ // S332
		Accept: 141,
		Ignore: "",
	},
	ActionRow{ // S333
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S334
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S335
		Accept: 113,
		Ignore: "",
	},
}
