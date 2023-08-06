package eckcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
)

func Test_Sign_Verify(t *testing.T) {
	testSignVerify(t, testCase_TTAK)
}

var (
	UserProvidedRandomInput = internal.HB(`
		73616C64 6A666177 70333939 75333734 72303938 7539385E 255E2568 6B72676E
		3B6C776B 72703437 74393363 25243839 34333938 35396B6A 646D6E76 636D2063
		766B206F 34753039 7220346A 206F6A32 6F757432 30397866 71773B6C 2A26215E
		23405523 2A232429 2823207A 20786F39 35377463 2D393520 35207635 6F697576
		39383736 20362076 6A206F35 6975762D 3035332C 6D63766C 726B6677 6F726574`)

	M = internal.HB(`
		54686973 20697320 61207361 6D706C65 206D6573 73616765 20666F72 2045432D
		4B434453 4120696D 706C656D 656E7461 74696F6E 2076616C 69646174 696F6E2E`)

	testCase_TTAK = []testCase{
		// Ⅱ.1 secp224r with SHA-224
		// p. 38
		{
			M:     M,
			curve: secp224r,
			hash:  hashSHA256_224,
			D:     internal.HI(`562A6F64 E162FFCB 51CD4707 774AE366 81B6CEF2 05FE5D43 912956A2`),
			Qx:    internal.HI(`B574169E 4FCEF1AF 3429D8BB 5481FF7D FA978690 492E1098 B80A5579`),
			Qy:    internal.HI(`1576819B D9F0B685 19EE844A FE88CCFB 2AD574A5 6472D954 1461AE7E`),
			K:     internal.HI(`76A0AFC1 8646D1B6 20A079FB 223865A7 BCB447F3 C03A35D8 78EA4CDA`),
			R:     internal.HI(`EEA58C91 E0CDCEB5 799B00D2 412D928F DD23122A 1C2BDF43 C2F8DAFA`),
			S:     internal.HI(`AEBAB53C 7A44A8B2 2F35FDB9 DE265F23 B89F65A6 9A8B7BD4 061911A6`),
		},
		// Ⅱ.2 secp224r with SHA-256
		// p. 40
		{
			M:     M,
			curve: secp224r,
			hash:  hashSHA256,
			D:     internal.HI(`61585827 449DBC0E C161B2CF 8575C9DF 149F41DD 0289BE4F F110773D`),
			Qx:    internal.HI(`CFA3A0C1 F4A0903E 84F314B0 70FB3EFA 531DB6D3 86739E01 2609557C`),
			Qy:    internal.HI(`FDB53330 B727A7B3 D40E332B 59AF060C 957D908D 18862159 F92B26B3`),
			K:     internal.HI(`EEC79D8D 4648DF3A 832A66E3 775537E0 00CC9B95 7E1319C5 DB9DD4F7`),
			R:     internal.HI(`64B49E97 7E6534F8 77CB68A3 806F6A98 9311CEAA 8A64A055 8077C04B`),
			S:     internal.HI(`AFF23D40 B1779511 51BE32F6 561B1B73 9E3E8F82 2CC52D4C B3909A93`),
		},
		// Ⅱ.3 secp256r with SHA-256
		// p. 42
		{
			M:     M,
			curve: secp256r,
			hash:  hashSHA256,
			D:     internal.HI(`9051A275 AA4D9843 9EDDED13 FA1C6CBB CCE775D8 CC9433DE E69C5984 8B3594DF`),
			Qx:    internal.HI(`148EDDD3 734FD5F1 5987579F 516089A8 C9FEF4AB 76B59D7B 8A01CDC5 6C4EDFDF`),
			Qy:    internal.HI(`A4E2E42C B4372A6F 2F3F71A1 49481549 F68D2963 539C853E 46B94696 569E8D61`),
			K:     internal.HI(`71B88F39 8916DA9C 90F555F1 B5732B7D C636B49C 638150BA C11BF05C FE16596A`),
			R:     internal.HI(`0EDDF680 601266EE 1DA83E55 A6D9445F C781DAEB 14C765E7 E5D0CDBA F1F14A68`),
			S:     internal.HI(`9B333457 661C7CF7 41BDDBC0 835553DF BB37EE74 F53DB699 E0A17780 C7B6F1D0`),
		},
		/**
		// Ⅱ.4 sect233r with SHA-224
		// p. 45
		{
			M:     M,
			curve: sect233r, // sect233r, Also known as: B-233, wap-wsg-idm-ecid-wtls11, ansit233r1
			hash:  hashSHA256_224,
			D:     internal.HI(`00BF 83825505 3DBF499C BE190DE3 5BC14AFC 1EA142F3 5EE69838 5B48D688`),
			Qx:    internal.HI(`01F4 85A65E59 B336E140 1C8A311F 01C92626 C663E69F 12A627E5 3E8F0675`),
			Qy:    internal.HI(`01BF 338CE75A DFB07DEB D962E1D8 0C101587 269AC995 1B40422B 12E9DA3E`),
			K:     internal.HI(`00F4 F088192E 8EB1CD8B 4ECB3A53 33746B40 EBF16966 A213B18A 176B2F62`),
			R:     internal.HI(`     82EF9427 4AC70A3D AC231E38 AE0F0D31 8FD8E189 EE40A3E0 61EC80BF`),
			S:     internal.HI(`00A8 CD7F7573 BAC3C4C4 00F65FDC CCD46F58 EBFC54CE 45571075 FD7704DB`),
		},
		// Ⅱ.5 sect233r with SHA-256
		// p. 47
		{
			M:     M,
			curve: sect233r,
			hash:  hashSHA256,
			D:     internal.HI(`000E D21C5C28 5F2B454A 0FE5D97C 6A86AA3F 7CB14FFD D35EB089 BE11F031`),
			Qx:    internal.HI(`0068 A50FAF91 43203612 1C5B6D2C 9307EA20 1FFE6F74 E09EF223 0AEC930F`),
			Qy:    internal.HI(`0052 67430AF3 EF4FB190 A4430F26 9521D7FC E007E221 245F5D14 C7541963`),
			K:     internal.HI(`0000 A516B3AD 24EB7F85 4D101DDF AB5A09A9 9D2C566A 09B29E57 2DAFDE75`),
			R:     internal.HI(`E1C9 75FBD0E8 98FDB018 61C4EC8D 4CEAE19B 8CFCBBC8 09EF3A03 AD3A853A`),
			S:     internal.HI(`00EF 7A6CAA2B 46D5BE07 DB837779 49F2505C 877FC475 76A54D40 BCD53D5F`),
		},
		// Ⅱ.6 sect233k with SHA-224
		// p. 49
		{
			M:     M,
			curve: sect233k,
			hash:  hashSHA256_224,
			D:     internal.HI(`0073 6439374F 72B1C723 AE611CB3 DFBCA0A8 E2C5096B DB9C2D37 21167B49`),
			Qx:    internal.HI(`01E9 1DEFBD41 AE655105 E046E03E C13E3860 0E9A2C9A 920B8E75 53721605`),
			Qy:    internal.HI(`0112 9C2706D1 9D134891 C7BAD84A 5600C2AF F86068C4 7497F5BD 498D0B76`),
			K:     internal.HI(`0061 7AA0B7A8 197A2B81 01500BFE 55D5322A 7149E275 F91ADBC7 E30128E4`),
			R:     internal.HI(`     B164A12F 615CC661 C10B78CB 6E01C9DE 46337C50 C036FAC5 51178752`),
			S:     internal.HI(`  4A 2109081E B3ADF95C 19FFAE89 5D303B83 147B27C6 EFAE8536 2BFAB89A`),
		},
		// Ⅱ.7 sect233k with SHA-256
		// p. 52
		{
			M:     M,
			curve: sect233k,
			hash:  hashSHA256,
			D:     internal.HI(`0028 13E18571 44BE8611 A7B93256 4EB3603D B08406A6 90D8185D EFE6EC5F`),
			Qx:    internal.HI(`00CF 66347977 26F04185 BC953B3D DB4B9375 9D074522 0938DA29 DD7FA585`),
			Qy:    internal.HI(`01E0 00F206CE D0896589 2BCFA3E7 B459CED4 7188EBDA 7A74B03E 2ECB66FC`),
			K:     internal.HI(`0006 B64D5DB8 65FD5E32 72ABFDBA EF0964F1 C26546A0 1582A4AD E5C0C2CF`),
			R:     internal.HI(`  D4 B2C6E695 9906C6A6 A8290AEF 7261FE96 EADCC177 63A1DE9D D009737C`),
			S:     internal.HI(`  08 657B1F0C DFCDF279 A6433A8D 68BA2D02 244A4C34 8519A05F A2CF37ED`),
		},
		// Ⅱ.8 sect283r with SHA-256
		// p. 54
		{
			M:     M,
			curve: sect283r,
			hash:  hashSHA256,
			D:     internal.HI(`00D64BEC 51F1ADA0 5BBD4F2B 53405B0C E8A1B99C D8DB6309 76A47F76 F08F205E EFC3FBD8`),
			Qx:    internal.HI(`04313C7E 9C4F80D2 6A287B37 FE7FAA96 BE31F116 2E18BDB4 70CF43D4 DB28DE10 8B007E9F`),
			Qy:    internal.HI(`0342CCF6 F502F9DF EC208170 24326C26 E867E1FB EC6634CB 17023CA0 222D6112 E0BFA106`),
			K:     internal.HI(`00D18E44 CB7F75F8 01277FA5 CF31A268 8CC2F322 2FA9F26E E8598126 AFEEE4E3 8DD0E08E`),
			R:     internal.HI(`         4A23BA73 B29A9010 ACD1E231 3B9A252C E209C7BF 3643926F A7BF8C87 A8C76D40`),
			S:     internal.HI(`03AA4FFF F1F4C3EE BF9C8798 2E717572 71CB7662 BA03463B 8B5F97B0 5C7F7C2C 88A31799`),
		},
		// Ⅱ.9 sect283r with SHA-256
		// p. 52
		{
			M:     M,
			curve: sect283r,
			hash:  hashSHA256,
			D:     internal.HI(`014930E6 6B51F09F EEBBAFFC 9111C5CF 8AE406C9 35AC9618 F0A613B9 6D97F7DB 8F6EBA74`),
			Qx:    internal.HI(`078A6ACD D5F779F2 5E8AB413 965E217F E6B1E63D 4717EEF5 0DC8C59D F7B1A095 BC3027AE`),
			Qy:    internal.HI(`07B6D962 5F2D9DDF 516B5037 E1E7B115 26E12AC4 E65AD498 CD85D65A 9E915D58 6976C00F`),
			K:     internal.HI(`01EA8FB5 72B7B2DA 7149DCD8 78101ECF 3F296400 E13A0D65 C8B6E558 C0237C6D A55268A1`),
			R:     internal.HI(`         E214F3CF 8BBB6E92 F779E6C8 A3424BA8 64734002 5EB49EED C6016746 81B14AFD`),
			S:     internal.HI(`0014CC0B B9245B7A 8BC3C6E0 392AAACE DCED8A61 9D9676E9 73D5244D 7F45E01D B425A93E`),
		},
		*/
	}
)
