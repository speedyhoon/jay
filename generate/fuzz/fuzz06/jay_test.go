package fuzz06

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA{}, expected)
	require.Equal(t, W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA{}, actual)

	actual = W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA{
		F1Oi0yLLnyFy062S3yjF0u8b8V765l504p0uO2: rando.Uint(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA{}, expected)
	require.NotEqual(t, W0rkOE1IkyT3uxkJm7j1A0uicgE06WXWL7GyJdUa7UF4Q5SbLew0AOSi25rtjQA{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_1(t *testing.T) {
	var expected, actual IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie{}, expected)
	require.Equal(t, IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie{}, actual)

	actual = IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie{
		EEiVJb38S8p5X253kt4ghHG021T4W6666bPs0A3JfQ12822wiKI3exxmCqx68exGt1oi6457b6QwiWXoOWHx78a1H:  rando.Int8(),
		LT64g0KW2hSaV2I5SbFoggp3KL7v8t3c8w3B15xqYo4wdJ7ravD42uvFqa7hwbenv2NL6hWHWs15R5wWfaNM8a6hy7: rando.Uint32(),
		HPy6ITP6hD3eiobQjyqHlBj7QitdkAj505SdpmFvT5I5:                                               struct{}{},
		JNO4goth: rando.Float64(),
		D0Um6KF2882kwOn2yI5k1EgOFAwgnQAp7k4mv2D6rSwQ8xOB3Rf6eKs81x5p71OM2Gl56W713l5CkW21j6YeHAdlcHwO8sGnL8N: rando.Uint(),
		K3S03Bj35GdKG1e4nS: rando.Uint8(),
		QsU1442N030PaQ11kd6Mgi4B84LPmIc2I6FpbU76n4R0Wob6bot6K7im1n1vDj2epM26HMU11D30URh8: rando.Bool(),
		B2u50f3QHqg7R:                       rando.Byte(),
		ReP40ULKA23qdehNyt266A81:            struct{}{},
		QaePCdF1ic0vQ42oClEUb266Mr8k0hoEcu4: rando.Time(),
		MaQWI3PqtQREd0HaSqwOXaMnQG:          rando.Uint16(),
		OgjUQm03PaKenDN5vD32wHA1UtObOk6s3i0u884Qx1NBr1jFMjLCLJ6bSKmF4Y:                               rando.Bytes(),
		UGR08sIK86oB32I7FEFTT6MC8gq0j1nYy6K2uLQES6T34O20K7W7QLp3:                                     rando.Int32(),
		Al55tk7rJBUeHmRh615riv8ui6od163vT8:                                                           rando.Float64(),
		RAp84uTFI514B0D620xsd6lxA6IFJ3gql6RR18MK3kym24:                                               rando.Int32(),
		SeL33jX568K4m11uSHYwyFVH2sPGCnojH3pde7ErsGU8D81831jRF3oa75I4SkiJ2I0FHGeI:                     rando.Int8(),
		J778EI4ils3L6kMO743k3xAG7NKI8x131G07N77kE6QUSdYM4El10tw1ivocky2mRfmbh0q8rU72N1NPp0LWivyh6KnM: rando.Uint64(),
		O88gA4FQG1RovDOHH2R8M77L1SDaJw488RH5F56K184P:                                                 rando.Bool(),
		DvHi8Bg5rb8RD85F18ca7T83WHouc2a2Y5b738U728eI1ceyHyRg71583f5j8o7i:                             rando.Int(),
		MKF7uUjp3gC1x6W5a1fsN46db1OL1Ly20Q2TmU48ipUR6llEi68sqVKDpA27F7H7p0w:                          rando.Int8(),
		HLI2KvIf586NsD7W1m6Q6I5nJXWW1VK76mbHo5NV1Sp2704RRNm0F:                                        rando.Float64(),
		R5shfjE6F3dbNgUGBeC3Wlq472N8VGyOIH7c4tJ6nun4dqPU6n65611B0WCKW6lrs2HCe:                        rando.Uint32(),
		NjFW2UyaoF15368BAQ4: rando.Int8(),
		D1oC21O65Kr6qxIgIa4NRXJBcSl8u8N0k1sN5IJmuc21J5LH2rsn0I7a5Ur2O086MrgotlY6x28TJtK00b5N605sbya8RM0pDl: rando.Rune(),
		FyYW8DQ5nkkTm6Lwn818U37k0nwDq30a3wu3Q36n38XR1wLw21f581858aA:                                        rando.Bytes(),
		KFpF1E315N7snJX2XFdDcXBioQfJ56f08MpTCdPbF7ll812e11T840hC327HO6DwR7gX38H2pquQ:                       rando.Float64(),
		J0gc5iCNd80C8wUE5iC5t3DaQdgaE: rando.Bool(),
		Ljq5ng8W66U2PHU3Yb:            rando.Float64(),
		JhHMOu7L5C7d6i447COeaqRQ26267plni7XA02lI0a7C5yle184S3CIF7f6D05Jt54FIk3Y5376KOGCfYL1472b: rando.Int16(),
		Lce8KXyL5ccuMdd4so74wwluAL3x5hOX11AJcRJbOX7I46TG07L1ggrAj6lU6:                           rando.Uint8(),
		J8i1BUrQKx5Tid107C33gV2g: rando.Int64(),
		HxIX7m1:                  struct{}{},
		M1118RenfXHc32R4Sh316IX3BNVT3o7vE2DNFj5QLkkF44TXY66r3416g4v07f:                      rando.String(),
		Ie4psGRIRsQdhh0xhxi6m4lyvY85oO:                                                      rando.Float64(),
		Oi6JO1WS3YLfwN2gYBvWVaH4ycK6aR852r5a5LytY0TQTYb8f1V7B4r1Dtx8qWv0a6L:                 rando.String(),
		NWqXk874i3Q5Ni2MfkF0dTO38OL31G5PHtOejNFSObeT5hg71jusX2n1TNgQ7vi2S4wr6A4DcTj:         rando.String(),
		Wyq2XX5SD1Xq0068pnwHMR45uRL61xnH5S35MVYujIfryhCgs4hGkIHI5ty83QnkEMH:                 rando.Uint32(),
		P6aX18ubRwqp7sIk3BBfB3I47s52FfPV6ogc4IbL1UuDwo5V3OBWR2NDj774b2VEnNqOCr2eY0i1711k7Um: rando.Float32(),
		Vb5XibpTWr8M8uy7:                                      rando.Uint16(),
		COjiSYSY5GcJ621QBTQoajAyPfn16C5ikbud:                  rando.Int8(),
		X3wuiG3pf6xP4ALWK5paR7tIxjn:                           rando.Byte(),
		F2MRs042vx5u8sh64k513V206g52H20FePq20px43uKCP6wGqsrvE: rando.Int16(),
		F0TvF5vdL0Uf653wKkBdVsOikgfe5o6yQ430CA665js878B0Q3c7wLG8eB10O7FVRug0OLYC45K72IV3: rando.Time(),
		N6hduYNKiAGgw8He58Vt54J8M2s1D:                                                                 rando.Bool(),
		Wv0OIIY368rWjghi5h122CF2B7vMmMH8Dn7vXeMAxhoM05t3r:                                             rando.Int8(),
		RkisGHlVM12UtYh4GfS8BVqyvf0kDYT08kM07RW:                                                       rando.Byte(),
		HYp46D2UNc03uq4DyILx3VIC0o32IkykQO1yl7snKEbkSD0rJ5Wci:                                         rando.Int64(),
		Hyhh82H1cGfWt4XkTJb03U6o144Pk13G26K4J83lFSTX8Li3U76CLHAps:                                     rando.Time(),
		L3aRJI3V43d1gsUKC81HXAf:                                                                       rando.Bool(),
		QXt844KDsMhE8Vq3vEmmwmY:                                                                       rando.Int32(),
		YkR54bcJ5b3xm0A5fAcomxC1n67I4d346380gKHChhQ52qd1m3A87vKGcSTOp:                                 struct{}{},
		L7k3QWB1siwuOuFyC7Iy8X0ypH2wacycnv6NvqDX051mUbSIsnsj:                                          rando.Int8(),
		Q4FmG6BU8Fet18ylpT41m8aUd2j16w2EWR7:                                                           rando.Int(),
		I6DRHaVh4WaW510NppxSX8DD:                                                                      rando.Int(),
		Bifxe4W3alSUX3K2RH38O08DR73sS0885w0O27Q6m7:                                                    rando.Uint64(),
		SUyS1oCgS8j4Njvc52a3j6o424YY567dJIvkh5UT:                                                      rando.Int(),
		Ih2d7440gp3pv2S6j6n1y0tr3Kn1m6Jg1vdPSYYv72qg23OQ86yTSBu5WTChCKgp8kI:                           rando.Uint16(),
		I3Hvjv38T8vRp54qJUSCvwemIwf3UtO6L0767310mJ73sk5mg0B2f13PDFq7vnY240e5bN8dEUx7utROW8Y1P:         rando.Int32(),
		K3JVrr228n73c7tOVSlU6DlJV12W4oR3aV1NLCjG8373D0Kja2ad13CY1E1cm2l5S1w4IUCCmljG86JJ53W3f167O3j6f: rando.Int16(),
		LQ1oJ75Fyro5kQwoqr4sr0VsW2KpG6kruu5yx78aQEMAG66hch1eEy3Iio:                                    rando.Uint32(),
		Dci7Kn3dgC1tkk0k0KLsY5x7yE147h1:                                                               rando.Uint8(),
		PeE26mhLd2J:                                                                                   struct{}{},
		Hvm4P06OXL6k3o4631Er7wU54Tf8420Jd7LdJ3YWRJ6wj5W03UP40WEG7VA4bABNuY2eY:                         rando.Int8(),
		T51817BB65R2S72t6X86rL205ML2xyMA10m4PpfXbNuFPbosI25Q1:                                         rando.Rune(),
		B:        rando.Int16(),
		Xe87tm80: rando.Int16(),
		P5e3W4evH50IbGq3nRTNK0Il01Q5Dc5831B6VwkjjPifga0L641o84ct27UOAWq1d3372v1n8ugmmc8eKjJjU2: rando.Bool(),
		P4x3e: rando.Time(),
		VemOJLPP576KlAO0G5Q8NtUXxUD2XaIWHU6fEYjs7u83xIRN4l: rando.Int64(),
		Nkp7025aXM8NUQkVP:   rando.String(),
		Eh3jDv127doM1uMrRa7: struct{}{},
		G1L2by6tT2odPm32NO:  rando.Bool(),
		IxTV3nE6566liq:      struct{}{},
		Q5lcvqUst40jiGKEdh2WTCtVP8yvD1pxnPyQNEu4eqKU22p8p66qim8l8sD48xJsCV75CKaU4Sbl1B11JGWy5ocdv6K61WJ: rando.Uint8(),
		HCT8Nm4tgw5E7cR6417:                                 rando.Time(),
		Sy1D48soMeWp1qJ5KQ0qglKf5LDX867CNUeWk3rOSO:          rando.Uint32(),
		Ov8Y0PTrv78Q05DfMRJ7gayi55cKbd084UfYuK87snwp7fBOfCk: rando.Uint8(),
		Hom32vQ2QShB2N048J8IOH771172akJLPnCrFc3GhLmF85VkY:   rando.Int(),
		StH4DfmY5sBC1vR308e23d0gu7j5K1sL2Dhdb51hc1f:         rando.Uint64(),
		DfM6H7JeYOAMn6L7:                                    struct{}{},
		PkBN72ctrToNr88tJ4yuno76:                            rando.Float64(),
		A65GXrpy22Xp8sF0eH54fB1:                             rando.Byte(),
		Va:                                                  rando.Int16(),
		P316f1l105Vt1Msorsau4s056J4qfcYsGe8GPGukGMSsEbCdR0Ou23a3XOktRbu68juSvv128VdJV1hvrMw2N:          rando.String(),
		Ak18R7F8R8Xi234006NHmeT6cPx64hMAaAoEuCcd0C2TxTW6nJ3C15BKW0d0tcXALSm06UHy6v1B23V210x3mwj2v8svpQ: rando.Uint(),
		Xtrkkn67FIydxvnyGnR5086R7YiY4vNQI: rando.Int8(),
		CI05Wj25b7Dv5LMkB825N70IfI671:     rando.Float32(),
		YyRH3C30TjI3at1h33p43M8s0T7tGx0d55hW2WcSXSAjCTkNmcHfi7aP6BXkNrK4eKLeOVHCc81YM40l0DNX53GhDuf27Y5: rando.Rune(),
		T00I3b7lX8iyi235OScjpXVod28xDCFgcsdNvMt55:                                                       rando.Rune(),
		SpDJbv4SOKhqgc2ek61NWCG13GOxL8ce75hNrRnhhRK674M278v8IsNRQPvoJ62FrSWheA31K8apiLsjh5C42cwOgU:      rando.Uint(),
		UvIs:            rando.Uint8(),
		OV11JU23EFV02jf: rando.Bool(),
		B54dOh2fldnFtih23FMak4elL26L80UEtKSC3RNt441JYm82D1uYcGKo8ysBY0:                         rando.Rune(),
		Ej58j8VwQlUj51i43gGrrNYDgXf53dR5QsNeR43u018x1OUYpYyB3G50:                               rando.Rune(),
		Jemrk4Y55ED66CSa0Ka1Th58Sp4kqb50ej83S3U5KJfbU4MG5nK6Y47o:                               rando.Uint8(),
		Bq73L2sA6tg7241rq72gH5pME:                                                              rando.Uint16(),
		TS7sG4e32fKI5G53XP7fu4JOkHrn1dpWr2yCmIYWh4J:                                            rando.Time(),
		C0u7ubl46R6R8A5k0220X6lO:                                                               rando.Int16(),
		L7TGNQ1cGDu4knWRK87p2sD0715N2A6tl526s68pvrO7pK7MBVxnxC5G5l1HDAg0GL3o8arUKN657FVs733D8:  rando.Int64(),
		QruMd754CMY70VDefqimgYxyK32fsSeKJYm8M04JNF5D0w28yGGc6wt7W11gercKI7XQ66N20Xw6NBJk5WI8XF: struct{}{},
		L08CE0y7ljFcoK1mQtR02b24p8pL37sefjM:                                                    struct{}{},
		YeCflxH7OtDL3n51ond1n8ce22K3QB8oKiI40Gie5xw2dqew78i:                                    rando.Time(),
		XWDgd8JCe6334I5rSFt5578s3ssyq3sJXxsB8D8x5107vuim47yeYWr7bL6bO6Pr6Fq5v7v17Pv4bWE:        rando.Uint(),
		VPc:             rando.Bool(),
		D268CqMHyonIDvD: rando.Int64(),
		Wxm62k7mugqT32mSO58C45H3Pe1fL5BJi7d4a6QCuypb3ufefh0TN0R222Fvg4G074l:     rando.Int16(),
		SS0d0pOseK0ci1q8QnV3OkuObvY47n713buR8y:                                  rando.Int(),
		C3hq0XmO4g4X8uo750Fuc8e4hBdIqCBkDEp83CFT6i8NIxcT5IQBSoBuXA0A0bpam3aUMxS: rando.Int32(),
		Lu2u0d4f5l: rando.Time(),
		TwjGrSi50GchNnPiJjDO2ppI70U2C5uKKmOT0N6EWyY5gdD3XUqQ64SUK4PDI06adbEdkb75wU: rando.Uint16(),
		G4XrBn28r7ur85F7WH5JY1: rando.Uint32(),
		S6:                     rando.Uint8(),
		P16x5p0I5C4F13tgdOx1NMaFk3p58LHV35jnx33NMoSf7MU841Lm:                                               rando.Float64(),
		Q1RQPL88Di0yJ42vDmgY1J5UGg8jHu4B11I06s4Q65CJ148N123VmcE033J7ets30C3XQlFAmT63LbWGS4mX28D5:           rando.Int(),
		STGUQco40wvbI1N0Kc60S81oVk174HBu6Cqm553w:                                                           rando.Uint8(),
		G7mAnPFV3XCQEsXPNfuiY3tyakY43AL6U6n30W:                                                             rando.Byte(),
		VPomO8URx1r1au07rDGO6uRlaaFPbKy2:                                                                   rando.Bytes(),
		A2tFC746uoADn32s8nGT2G3D56AqXulgLL02bryhwrql438ImTD2l4D4feTnOcAD2u40:                               rando.Int16(),
		A86xfjhfEK8t8J2nLvi76pLBO6:                                                                         rando.Uint32(),
		U5k2Y1U175X3gN5PCChGpK5Mnl25Syg8LxQ3DC01Ob6GvtSg631rdfn:                                            struct{}{},
		AH4Lnyg5K2dWwqJnY4nK8UiSc2LyH7TvXNu5pQLddsA5:                                                       rando.Bytes(),
		WsfuNFWGfQ2K1UijWSB3vt4yT4aOlNy3oy7EttvjQo2v62:                                                     rando.Int16(),
		OUgoD26C601RW4Po4y23KyU82CfEIBgHFW57mh3c:                                                           rando.Int16(),
		T86D1bY7hyQ3uPO7paCtVt6q2djrt5dNGGwpK1OrK5U3JvRH1waLdgk118Po8mMu3gWmuPVmO3Q4G7B2Dtv7uE1wyll63B23CS: rando.Int32(),
		QYFSQBb8:                    rando.Bool(),
		CRWJK25cLfl:                 rando.Float32(),
		Qr5M7E7Vii12A0oeTS:          rando.Byte(),
		Y6Si04J44UAJ0u53AXUWKSr032u: rando.Uint64(),
		M8j5e2Q4xSh00xd3kd36cP:      rando.Int(),
		CdncJBk5WD4liJSj1TdpH481wF8CkHc6mpMinfV7juuvg0aBc3r606oRP46h56vO0E0:                               rando.Uint64(),
		BuORF4L4WV7lV1jaXBi4vPrVXATVx1Gd7e6ywiw4Xjx1v8411FeGy4x5WwG3hKob4mGKK24J1au4CQ3SnRBJH4JW:          rando.Bytes(),
		PLPPbc2Ie8O8rg53Tiv5W5CqFsG3YPXT8d8cr0Glt4cUw8lo8iDuGBxIMrvC1uHH5wNtVg825H15aCBkBWvT3GXXSw1p0QjT0: rando.Uint(),
		UH4xPp1ul5e7q12: rando.Bytes(),
		Q7ud664857WOoi2Y6Hd3Px8w02A0wlP5hFO5rw7c4J1Y3BH3PyMLoC: rando.Bytes(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie{}, expected)
	require.NotEqual(t, IEm60mY4B4GoD52333Ot6a1Dk1s2S2ie{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N{}, expected)
	require.Equal(t, QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N{}, actual)

	actual = QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N{
		PpBFn602770rY615345kRiweiXo2oBn8xyflN3IkM80672: rando.Uint64(),
		Bj0k1H0Sb:             rando.Uint(),
		SirR:                  rando.Uint16(),
		KtiWmuWQjQ8QGKD4X82lU: struct{}{},
		MbpAxkR0qC41:          rando.Byte(),
		TLFh8QL56AO4MI7FGw1VEjCs1S43rK0Qwh6Rt3s0kA0EWR85: rando.Int8(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N{}, expected)
	require.NotEqual(t, QU3Ymj8o3j80P1YySWyVrOB50hOHv63U3pc2KUqkt7fvbh3N{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7{}, expected)
	require.Equal(t, YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7{}, actual)

	actual = YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7{
		G5V6BFov1B6nC7ND8SqikkF6xDr6ypUf0aMe50HO44Q1xC2P8XUg80CcR871qd4lK2TToh7356f5Fa8D4O86Uc06e3H2ejsD: rando.Bytes(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7{}, expected)
	require.NotEqual(t, YGE1U7snuv8fY6jvYGBRyIWQRMFeb6F4RQUuJUTHyq7a24Ft153kGK70i35m2L7XJaXS5eWCxmASOs7Vk7{}, actual)
	require.Equal(t, expected, actual)
}
