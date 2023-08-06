package eckcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
)

func Test_Verify_P256_SHA224(t *testing.T) {
	testVerify(t, testCases_P256_SHA224)
}

var testCases_P256_SHA224 = []testCase{
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-256)(SHA-224)_SGT.txt
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`A329BE071A7DD3D9A349004C5FEF531BD805E2F3CAA115290C39170CC0B4D291769E609A68205AB56FA8E97EF0E604874509C600F4DF3ADF32CF91F301322474702A288A3C5F7C766EC985B324F370C926DBC32A7ADAB9D991B4341D444CC3173C089F76F464FB13FDD1A2C6AA4B3B2FCACD29CA9D0E5B00AAA20F999B09D73F`),
		Qx:    internal.HI(`B61A2B2E8D38CF62E0B42A3F7DFCEC28E6551DAD31334CDFEB67D4474AD60006`),
		Qy:    internal.HI(`08BDFACC750E0C7EB7DF6002F0651216F79B1727C152A060A7953E4EFA42F30A`),
		R:     internal.HI(`58DF18A4B633E721C348220B8BFC46C155DCABCC5C4B579C16CB8614`),
		S:     internal.HI(`29B5C07C3980197536E4C98207BBDA41E45BC12C8DC3A51D2D13AD1B54911638`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`CE9CD868B400A369C20F92F9F9C914F3C957CD9B66A0EE26477EBBDEE219929761746ED26AAD483FACFA33F750F0AF34021A0DC3A5604B077A5059CBBBD22DB37A44838CCA8CC442EE30E8D4B6ECCAFB9DDED1CFF4C150429C4C91A79136E910E4E8C1D65067FB0909789E0BF8912A370B9486B3B4E2E46450CC922296EA1011`),
		Qx:    internal.HI(`7F1AB01C0D368CE8211F9F6BB17FF892E2AEA1A02900E89F93EEE31630506363`),
		Qy:    internal.HI(`E4D3A8AF7074310CB56BFB3DB5647C835DB3805181E240E701E1FB765A27F00C`),
		R:     internal.HI(`0A074E8BFE24EBEFE8697C41704C05A8B722743EFD60C24565CFD10B`),
		S:     internal.HI(`C7908FB729730EB29998D1E5763F50ED2122C8857C3792647DC8CA670EBA9D66`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`B4655A5C43AFD911C1C0C97FAF599CC0035422CD0C88F7DD7F521751C640F30344E8092CAD7E50C12420667B396AF657855F24E2EF2181085B8A5DB6E42EEB1939C510C165965B6888D5F9B1BD8D15A3D91D921F14C280275E462F3579535A707F7939FB0860003A8A468E3E2937DE368E1FFB19FC29FC644A80DC1FDAF9A80B`),
		Qx:    internal.HI(`2DD7C8D08540128EC91365061E183559DFB93C8519551E41EBC4555A3BAB9D44`),
		Qy:    internal.HI(`040982DC6E77F4363E2930DA905D4389345BC1BABCA8740E87C23E1C0563B4B7`),
		R:     internal.HI(`52E33C53DF6BB66DBE7B30486EE08CE25684028ED00C8D90D9004049`),
		S:     internal.HI(`01C31F77643409DD0717DEE3B51F05056602804A9DCE951AE43C8E76D057C06C`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`49A8D724F9EE47561201BAA335843F867C1E75CDF1BB04D0295B3D2B200D48D8141FC4EC66F946C24BBE98CD5F362CF3C3FA9EA00387B0674B21B2772F28B3A9A2D2636D43E2166CAD5BCD0EEEBAA4C7CCBC9C5F0C411E0A4C46238BB28A693B01DE9C284EB4DC2AF5DE8621F022AC3045941E40AB4777830D6301561A19F532`),
		Qx:    internal.HI(`7586F3F6CA5B30046FA5E2D036EC4BF072BB53167A873E4DFCC7105A7D185D7C`),
		Qy:    internal.HI(`DD0701776BB40F10370A59F69052C7F4C9B998F5489422B328390CA0F99E2F85`),
		R:     internal.HI(`89DCC6A6C1289FC8F716FDBAB962429EA2CDEC6619C6D45C3D4934F8`),
		S:     internal.HI(`75C0991F330FB8AFE246971CBD914B22105E255DCDA5CF0F9A42AF7CE1B17CE4`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`8288E4050B20C3BB297677293E2C504827DB5ADF499FE885B73C432FA465E71AC43C3457C87FFEC3977ADBB27439A60DB1110F4118F6ABA8BDBA6BD452A6D967A98F0FD598D4C8D3D46778AEFB56CC696BDF82D310A2FE70D9F1806EEEBC6C755C3C7EA158C7655EBEE69A780236E62A2515826BF2A128470C1A138A0C2E4988`),
		Qx:    internal.HI(`B4B0B6D3CE897A9083230F72193F138EDE02A8715D623E9F48EA9DD73A976730`),
		Qy:    internal.HI(`297B9021857A25AF512EAE05F488C100A155AF92D686338833F5088D0867ADC4`),
		R:     internal.HI(`7F00C545B787852FF9F248E8660BDF3758E54692AB717CC6FEABE363`),
		S:     internal.HI(`BB2DCE5327B86901543B07137FCA44B0C32B624A4BD0B6EBCFE0225232A35DE0`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`532A1544AEAA20C47CC314D47E35230AF7AD66464797797E9F993D20062B23CD4964EDB006774C4B7AF644ED2E57B7A941C60B0860D3474F26F89C90008CB157411EA83C96CF45206F9E0E559B46E08DAAAAD8BF554AF3DE79E95BA0E2CFB52287B772AB59FC6E595900DD081157E32823C8BBE00899E432BB48287F621CFA13`),
		Qx:    internal.HI(`6D2E68F6A27CE2C5B264390498048B4E28F09D53F0FB5BC9846DBF22C0959F87`),
		Qy:    internal.HI(`8C5E70CB2CC611240176988A250A064EAD100649D5C955A2D22829BB46FA4D61`),
		R:     internal.HI(`44501BF991C05F51B0408556FBDB7E9D8A926B1F0C08FA2B854EF72B`),
		S:     internal.HI(`BD13CE1083F7E742C9921BA6422DD42226DFB6C79914BC68719C621583280CE8`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`AFB1FD2515EF32F67D8BA769A9850DD2E2B92B472007884054163DC4F9444FF596BA823B564305DB69D7E7434174B4CA693F263A108058E0F87F5A6FEBBC8F7D5EA6C4E6723862D7F3A1A2C8806F34397E4032670F9BD2D6A0D4C7E642A69A4674720D8885B8CBA039D16595D469F52E31CF5DE11F957ECA8E9154F9D1C85AD6`),
		Qx:    internal.HI(`605DF8AE2BFEEE4BD5B371CAD21E35021B2473E9AED80CFACD28B1970A918FE9`),
		Qy:    internal.HI(`DCD18C522F711B47A8FD3F11CB5193419BD3FAF9FAB2410901B965271BF138CF`),
		R:     internal.HI(`F3CDB773E313AFDE433D06BBECF45A2F5673D480A70068F7FABD2053`),
		S:     internal.HI(`F0A45D2BFF025C5080236EC0E63303ABACB7BF74E72D5BAB0A421D08B3347BFC`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`8B4130EA7454CED4A17442AB740061A2DB233E260853EC4F4B5859DE3392C1979F62873CEB47FCF9D8C1D8776073F0751B9EF41B5B63B2DFA9F3B835CA1CC7DDF549F5186172F17DD4174ACB5EB31D6FD9C7240E71FB6FDDC355D804C2266DE61791E37F105E51B6D4FC44E2FC50703F4550FCB36CF7CC91F999A9BC0D15BED5`),
		Qx:    internal.HI(`0B519C399A22F210262A19918E4F7DBB42A6BF6E526AF1F81C98927E6E19CB36`),
		Qy:    internal.HI(`0DA4FBBB44636DD9E56C2F34E07ED8B04A6EBAC5A614BCDAB8438ECA26CFBAD9`),
		R:     internal.HI(`6062A2B1E07FD65C082600D08B4E322F575AE8E02673DF1894C5F7C7`),
		S:     internal.HI(`D26B2AD3803B3A83BEA59ADBC42F08A10C47791B000404F408B64B4631349E79`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`DA0043D9003DC6E35C1FFB5F9286747FD610342732DF772FF702A43266FCCBB55A8191F8FAE805283B572B4E3F3ABFAD4C0809EF76DF28D0ADF9CAA65090AC7CF92BCF1595E2C89586A21821E9F6EF33B06242F9B1CC9E775511A2BF16318405653986D12E53D2209C268FB43FEFA860516D2C992225A00C71053D8CCAE87B14`),
		Qx:    internal.HI(`C793C5E0370C73085C2D9D30DFB5641D2FCF40D122570503639A85A25B60A1E9`),
		Qy:    internal.HI(`62C10BC94D7E039193F90264327211088D4CA443533AE52F19C2B645D74D2207`),
		R:     internal.HI(`09071D1C5BCDE39D34A529C2B7D9E42DA68D512749A582BE41C2C66E`),
		S:     internal.HI(`A3C8B552568C6149A91FE49E44C7DBB6A22C4554B952F7EB3D6B7B70EEDE60AC`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`910FCA35EA0DF0A62133E449B800996CC7A29F8DD30FFE64CCB833844863C255B83B34B2B58AF4ED053EF38B93AB7678F1A2F9FA95588F377633A58631FB945D5F71E72144EBBAA47CE7218FD61EFD8AF735206C00743227CAAB3AD9F2AE31A7528D8DC414FB246105F35ACE512BF3954A4B82D97682CFC06978232DBC24E397`),
		Qx:    internal.HI(`C7C9F7465430053499BB4E8813541EE7D3F0197DBF91683A926A879346AEE9B4`),
		Qy:    internal.HI(`95DA1115EF3F7CC6F43078DE2712276A282155BEB119EA2E9D075E5B2286C2C8`),
		R:     internal.HI(`BB23C9DB90E9874909E5E155F002103092616921C83B0839B1FBCC38`),
		S:     internal.HI(`E2EDB9C9108F3D14908E340BBCB580BA20C818042400D2E1BD93E9C10B865C88`),
	},
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-256)(SHA-224)_SVT.txt
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`AEA3AF76D4C5E2126D9AE75AE15E1C861E1983BD943915C950E4F126DC064AD0F9A68888FE907DD6DBC6119E8CE258820BD273E2770626B8939686284466B3FB3ACCF9D665CD4778686863BCA908C2B12A916E8ED528DCC35A2927FD56618252DE076771188D4E0E13C0645F7A4C2AE1501F5B4563C40AB48B2F5E437806CAD9`),
		Qx:    internal.HI(`C2368943B66AA845F3EE022AD5836702A877F9F2E839253ABD0E20DFF4154730`),
		Qy:    internal.HI(`462BEAFAA5A9F02A4231600452ADA2FF7301A8EB3F100BA6806E993809BE1A8B`),
		R:     internal.HI(`511A7CA85E726877EBEE65155DE5CDC6F2E6E58C79E7FA7F59258E90`),
		S:     internal.HI(`F06A0DEB0A95FE25A34490652E84AA40926B2B451721326723D958548FF06745`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`F340506F092916A5B2A519884FE94A17EB2D751C270ABCD3E1BE80578B28FCF3D5F9550BA0F774B0CDBD64A5DDDF8BE10CDC7044C6BF73A621E422E94FC08C2B960F46D2DD5C748C444AA3ECE975AADA436A73673D92DA4278737F6CD83FC651CC07D4F56F89CD59F7F41EBD8B9A17F3F6E0A25850755CAB1815D667D5D84CBD`),
		Qx:    internal.HI(`AD78D0E2F6F1A23674CD81F6F9BC40302A264C76F23C422DE860D094F4F24F06`),
		Qy:    internal.HI(`D010AC7E359820E0D45E7D0752683B4D91EE52639A232E11AC68B2A93A2E8EE7`),
		R:     internal.HI(`FFD2B44AF974335253917D674746F664A36082686723E9C04C943DA0`),
		S:     internal.HI(`E6444408228E1E7064A58BF39F524D1900F1B435F5CE80DE8170A8AB85D249ED`),
		Fail:  true,
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`914F3F5A92E7B3280DAE092EC7664AE0407FD1AB37331735040C3C5FBF1E7D8988BB3527F3EC9C2ADF6CE24DC082D6BFAB6BBA4920AE745BEA11B926038E6A2FE7F6B57D409C3DD742CDD31DBA5508501FE01D439B81A08674223686B020DC8573ED304DE6EDF64B1D5D7EF9E78E113AE6410D9FDA57E0258BBC796F3AB66B06`),
		Qx:    internal.HI(`CB5952C2632BB0B004F739E79E40F6DBED067F68C1225904C6C11D983B032D20`),
		Qy:    internal.HI(`7914519AFA61BCBDA4C76B7D07863FD71A8A09EB6A731FA5EE50997492E0701C`),
		R:     internal.HI(`345275699BE335D7FD3A0D363200DB4D02489B3BE5025974C30B9779`),
		S:     internal.HI(`51265677DB428718F39446FCE4A525692C6A7F4873EF41BDFF0FD27F22B93939`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`7B5CB0DB3CF04C5E8CCB64FE45ECCD51DAE6889640405424468EF69E9D74D8505D3BE59E4E1B49FEB4EE0840F85D00B9E04C5412B696D33A0ECD5C60861FC5523E135536503255FB1960B8758EEE73806F2A7E18997E0CD04AD2058C8D5B40E8AC749230D7B42F5DE4E26EED57F26D0F8AB5ABFB335D11F2188481B736DE66FC`),
		Qx:    internal.HI(`73B5A7A59643BEAE63554B7899A654B581D39276EB8B7E10ECBCBFEAA7B479BB`),
		Qy:    internal.HI(`27821104BE40B3135DA2C5CD2CF3176F0F04D4205D08C6CC7C00645F910956BC`),
		R:     internal.HI(`2A773905176156B738C3805AF6209D9D2841F577FC0D15925A141797`),
		S:     internal.HI(`2A45815D70144BFD6E2E9B2D32FF9221EDA6F418881AB24189A05E27DCD0C78A`),
		Fail:  true,
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`C4697C14D95E600298906379C6FC264F4B2FC952A76F833736B3EC4A4CD28092EB0931B33159786325BAC6AD08C3C8CCE0AF59F1556B97A88E91F8393AF8CAC16ACA0E4A5BB009B1EF257F9DDF66F3526C13E5A69B0D66DDBB88F1BE50EAF09F107ADFAF17625956F80B1DD8E16D5025643F01FA1EA83B35E365AA9D840221E3`),
		Qx:    internal.HI(`0BD78E41F7B624F1D49B524500B64BA3BE6B453AB80FE68365C9B9D7A4B7FD3A`),
		Qy:    internal.HI(`C8F68DAC8D9D78AC7BC80A2AEB251F985BFA9B2E3FABB6A5635195DF4C367F65`),
		R:     internal.HI(`2A378AA596C6F86F2B0E53EAE6AA507A9B8B3AF258971C1E97FA5792`),
		S:     internal.HI(`EFF293924CB3D5915974656FC87AF66548ADEBFA3B8A49120CAF9E441A81137B`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`14FB1FED13B8A28ACEB95666089272823FC6DE4C346AF5742181019445912D7DB8423E30127DD237F9F474930D704CCA48D7AB59534D4C23D85A7C085E7E054BAE1819F6550706D93FF5FB52FEADAF89B75E234250484FEFD3A00A2EA02DDC50666C272277BD48917A6E185501117F52100D6DCB1AC6FE1E7F490D1D07EB09C4`),
		Qx:    internal.HI(`85429E25DE038B6FAA0870C72EF03A85B5D4228D9420C9B7AFC9B1A64F8EE5B4`),
		Qy:    internal.HI(`C9A7A2D32BE90059BE6F180D57663E67C12D8C83A897867418452C71163C3DBB`),
		R:     internal.HI(`9618472C5EBC97935FA4AD9A91F3E2C6B08C48F0ED91C595199158DB`),
		S:     internal.HI(`B1663332740223871C4FD4665F3D8EE28A15C2B26E44C02C31EC045ECC74F7E2`),
		Fail:  true,
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`3DC55ACA3C42CBD00E6A9079109ABD9169517E1FCA92064EED01628C4A77DC02D35717DD37708F7788AF9BE9A671A3056CAEB151E9D8FDF4AEA8A7347CA53E7E4CAF3F801353386E2B36AEDFC5517E64575DC64C5B1B1FB25FDA3543BBF176E0D680F390763625CEE76CBDC4766021711CC8C242CA99F135353A63F393F5FC03`),
		Qx:    internal.HI(`4FD2440F953AA3A593BD38B768DA2C5634711451DF2DE4AD1090D77F55E7A7C7`),
		Qy:    internal.HI(`2F7C43D5173A30512E5225D75A7CF16ACB28EEE0DFD88DAAA3E313435FC5F29E`),
		R:     internal.HI(`F0E2E1A29318A020F9B2599F8CAF64C11553AB85ACF9306767810556`),
		S:     internal.HI(`BCE8AD513254BD4EBA7EED2850378D859C5F613577D5DC03F1CC4EEE5A40B308`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`E91B930F4901DB362982A8CBA8A91938C0720DEB794355ECD34F289A2F8FE67D6E69B696A8BDFE16AC7F41DC1F7A71AA7E6B2E78C72910C919E16E6E0938B29EFD54CFCA5494CB119B00F0CC9F3EB500A30ACE357C018D97E7C7F8B3D245B7DAEA1F5B723807FB370718BCC3EEEA49473773C0EC73711AAB7F77A118C24F6E06`),
		Qx:    internal.HI(`A140C4C2043BB5777C52564FC972F31FC3BDB83D7E2C595D557FF09328878AC8`),
		Qy:    internal.HI(`EC37C05D7619CFD89205880E35889439245C1B227FCFE18A9C9419B2A3F9310A`),
		R:     internal.HI(`C3743CA68F6ACD0C830452CDEA99373DE2775A89730C11E7C0E03177`),
		S:     internal.HI(`8329DB4633444942CF36780F99A81DDE30593D3831A0FC73748B8CBE8F0BE102`),
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`D2815E25E0722166A155F2B7BABC1FAD947C7275672EE18C31EE51F93679813C3D0C65D6313AED973E11948F32C0A7DD2025A22D19679E22FC0B1C7F219E3153FA7466EC538645E15AF2E9F01D328A0896835F34058BCEBCA6C93E4581CF842D6AD649F3E9AF5AF3289D3B418F18FAD839099406F08BE53DC6A684230C247027`),
		Qx:    internal.HI(`FF03E76CA809C4FDD5A828975473A1E7853DCFFA6A974C1536545884B6A0E610`),
		Qy:    internal.HI(`CF7376979DBC4A066FB69F95DFC43579549E0BC3CF9B7CFFA8B19CF642539ABF`),
		R:     internal.HI(`C244BB9228C77805BFD1FF6A895AD08BC5508A121D4135E3B587E9B6`),
		S:     internal.HI(`795033F7479926C3448C4A424A4C71F523F57878D2A2BCBE9ED5BBF7D75D11E7`),
		Fail:  true,
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`58BF6C6AD5906A3AB9D9E91AE186132E6C04BD38E1FCC1A28C46D12C3BEB0C42C0A524BB62C44A1DA7DAC51798285A49EF186D1950C90BA600FEB4942E6627C850A0F1CCA6E51CF656B9AC064397455315F3BB42808BD1205357F80F51164D2767D79F7E3D66CF37AC00423A2546BDD18EC265B99FF836E7405FB494FD321BC7`),
		Qx:    internal.HI(`C2022F1BC5C619E694329219C819037E92B782CBE1E7B53460A2A7E279007990`),
		Qy:    internal.HI(`6AA906118D69BE1F61B16195A9E1B249B7240A7FAE8BE17A3C13FC788FF1ACD4`),
		R:     internal.HI(`8A4296BB71220D490F38A86B5A4E3613BC1F047DCB3CA5A7E676CD87`),
		S:     internal.HI(`E46895C3070414D82D327A278AE5C15F11D248937ED032CE6603DE8DA08086F3`),
		Fail:  true,
	},
	{
		curve: p256,
		hash:  hashSHA256_224,
		M:     internal.HB(`0037091AFD97FAC7624F7B1F40A1EE0E1E8B1786DEC05585C81F5F69B2E8C5189979CE0EACEAD8CB1834F2AE8F44CAD0EBB69BD1CF814635C164374D8B4A4835618241F0AB08212BD5A5D8B51E339C20B9CF76B038374400A629C1872C00CB65C05B8598AB1F89106585C4B7A09729F1F5D5ED920E4565E2C4C734AE0CE7F5ED`),
		Qx:    internal.HI(`6A6EE0930056C8CF1C2ED6655B46F257996CAF3A239D70B6FF946871DF0650F0`),
		Qy:    internal.HI(`BC13C09BC94BBEEED3D11AA9249871FA1A40E1DC99C79BD7B76BEB20D1A03613`),
		R:     internal.HI(`9FDE2AF04EA2ECA4E5AB607432CEA72F49C6CD9BA50B5EA6C9597D9F`),
		S:     internal.HI(`DF5B892E6FC0F43E206EF43F0E4D9C5C745EC7F4AF2D47B42B20F26DA6FA87BA`),
	},
}
