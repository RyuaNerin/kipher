package eckcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
)

func Test_Verify_P224_SHA224(t *testing.T) {
	testVerify(t, testCases_P224_SHA224)
}

var testCases_P224_SHA224 = []testCase{
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-224)(SHA-224)_SGT.txt
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`C4131E793A35F8C12465F71440AEA07759AEA79773A2F277F29E06BAD281575CA0A2581962C5BD510863CBB199C02763BDE91011DA0F5783515A1B2023AB9A1A8B5E0287CA98BA41BC1CF4B9693B069F36E05D9ADF4484DCDE551CCF734636BEE26EE79F20E292E38FE1C281F39B0414454C9F429449A2448DF21CD0E824F828`),
		Qx:    internal.HI(`B918A20CAEC26FCF1357B6E4D2863C86F8D50FE6E7E8A7E6498D73C1`),
		Qy:    internal.HI(`341EC41F68BD89669ACFDC9F4F036E5A25AB67191C33AE4D552AF9D5`),
		R:     internal.HI(`4FACC82F5AA50FB16CE07E37B7BB77F026BB147CCE816D98B55E01B3`),
		S:     internal.HI(`B61314B9CD93419B8D6651E8B2868A5C10EC755A180777A4393137EB`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`4B205CA4DB7DF36D192A0AC73F06EE94B4116E7ED4CBE36999BFB4946EBFB21EB5EA3E4039ED920139982AAF1D499EDA9539EB2F422480DE876F06F4BB383F4D546CF087425E2DA718F209C726DBF28842F6BE9A4A53BB2023E0BE314A5A844CA66C5FF3C5A48A4E252B1604BCD8D2275CA416701D00DC927F160DF561946B98`),
		Qx:    internal.HI(`FE6FF58FDE2597BB6AB31D1EAD9A09764F2F83BBD9ADBE1F92D51D82`),
		Qy:    internal.HI(`FA8307AECF867492144B9F9B2CB21F4388F979CE7CC0E6E40864A22A`),
		R:     internal.HI(`3BBE019CAFD41F9E0A908E17E52C271A1BEFB04D71592218B71C6853`),
		S:     internal.HI(`1EB61E1EB2E29744A8DF07010AAB641C6A5E5005BE55186819B72D19`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`72D80068573A74CE185946C207062D2BCDC4ADC8B6773581F93EAE500C02E799C5C7CD315943C2E8F1BF18C8F8D3FAEBC702C31FD9B4202033ABE2BB2164F49BF5784C01F08C5243F8C49241485538ACC42F1F5E1155FD8C1ADC76183C2496D5A0518AF30C870771BCF943C47748EF97864713D4AFA236C8E174B9AFC254D765`),
		Qx:    internal.HI(`F4EA15096DB067D78BC322AD347F8953421FF1DAB7C7D67F38E5B582`),
		Qy:    internal.HI(`AC643F0634671641040D73E9E1E6EA1967C939B339229BD19333E457`),
		R:     internal.HI(`73D74DD63F693B15DEDE91B51F20EBA14E154D11FA6BC35C66D882D5`),
		S:     internal.HI(`10EC295B17F0D0B62AD3F9C1AB701971BCB18D65D5F886AE11DA1F81`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`ABE01B8A6351D0E71517414CCF132FBE176C7737CF893CC1053F8832DFADCB5145DD1AB174ACA20824FD294060C20E1BC8E8ABA753A4894C483344BA89908C87E3252CBC87057A1B4FB6236B020DAB90322E97AAE7309E22B86CD8C77D0940DD43C07E64A96F5E4F4770DD04584F2FE739D9AB31FE1704EAA72FB6433FC71012`),
		Qx:    internal.HI(`25DADBDFF3FC33A6231A5D10AEC01C41F1E3546DBE675EE3456C5C1F`),
		Qy:    internal.HI(`ADD6736A4C24D9762BD12EEA8A0C9B0AAD0347237185F3A029FAAF73`),
		R:     internal.HI(`FBDA9BF9C285760E515EE4F14240A163AAE88DBDC45D6DBBFB44D5F8`),
		S:     internal.HI(`F721C9E0C9CD6AAD066536256C11832D5F5E104C903583C8173B3921`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`6BDCC3CEB2A65ABD04868DA8C98FC9D206ACE191D1E64C2FB1E7D87F1B2430CAA9CF3884410C8566C577F15A867BAFEC0A90B78A64D71165BB2AC0332823DC959218A279BBAFFB3013EB5089896721B60098374282C8F2E8F0B57982416C56E9055E4D095140E2EDBCB5780A9251669AE800F14ABD419AFDC46B96F30C51EA23`),
		Qx:    internal.HI(`6400C62EF64DF9D81C886700545E7515A640E6E2F434293356B07785`),
		Qy:    internal.HI(`355A6E4DD1692F5595B9A3C097FCFC02DD5F212FCB9DCF9EDBA65767`),
		R:     internal.HI(`7EB9A36C27FC58395450C31ED18DE0DB071C00EA90228190E3B0BC06`),
		S:     internal.HI(`F9C29957466CAEFE0353C6D49A73FC34525C4E181B4691A9A70550CC`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`266F0CF7F91C6754D9CBC0192AE0CEEB0E28FD987171B9CEF25A2F79F4CCEC8965423D6D7147C005CA4F055BA160B0E3039CFC8BC0310A7180B5E96C307FB64A76F5C300416C29893689AEDF11C76BA3A00F15EB95004DE2B7DAED8DBCB2AC7C58CF0CA7B7DEE84F0EEAA9185AB36736085CF9E5A2044C032C4DED065D583A1D`),
		Qx:    internal.HI(`0859AAC394F70A38A945C725E6D78BE5B458D2D26C8FD6F9D63F93D7`),
		Qy:    internal.HI(`85E62C38CA14790AF153AB19823657AA7EC85A7B7A7E72B77E846A02`),
		R:     internal.HI(`0E79E54EF22BE1AF9BFD89A75D93FFC44727638A0C52052076698FDB`),
		S:     internal.HI(`45D0556C501C12AC0CB2AC76F1D21399D1EF9D93F0700DAB04990CEC`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`503F0ACAEB994AAF880A6DE42569138DA484E112630FD7A3BCBB24669F09D212ECDA3B31BA41A5EA26AAF787E3D6E68326B28D701A96C97389F754A8D708F02A0260A310CC215728AEB3CFB0CD9160DB88394467D3BB03120000C72D233F171BB1B7CE02902CC37831340273E3D7073D0C94D8C560456D02D5F950BD663ED383`),
		Qx:    internal.HI(`4D8991D0E0A63821B08CB1D58471F64EBFB7B81EC42E26E4CF7E1094`),
		Qy:    internal.HI(`B7E1AD9A36DD7C2A4D759D28A38F8C1C1DBC8BD5D8363922BEE53063`),
		R:     internal.HI(`EEDB4BBB69579AA28EBAFC384A975C9BCB638019C3D3986456031BE8`),
		S:     internal.HI(`56EBC155539A0EE23F1BF5127F99480CEFA9D09FBCAE5F27E4A8CD4D`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`5CEED10A3C0058D30667274DEF8E6C3C3B64A1C25BA2FAB0022E498A4E3FB6E8B33B4793CFDD8B19CDAD5D2281422552E7757EFC28EBA270CD15942C50245DB8ACFC567011B2D9126E8C4841F22AD2E22BB9D97BF1E0687DC04A9CA5A8766A4983B9A9DD8F10C86D18B8195F62231A34674BA1AEAAE753FDB091525E5B6789D9`),
		Qx:    internal.HI(`DFAFC30E26D0C4BBDFD69DFB2706B88B202A4A77CF0E6FF70471A716`),
		Qy:    internal.HI(`FE75DD6A0F44032946EBA8ACF7B6625E834C67337E5234316C751DA0`),
		R:     internal.HI(`F2D19926E8E4B7FC4C86F0903DC506B75B1FF2EF5764C6DECD84C8A2`),
		S:     internal.HI(`7ADB126AB07FE741278FFDD8589A26ADF643E47047AED653BBA0C22C`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`BF20757CA134E4C445068497BCB3AD7C476D4F6B0D1076FAB9D8322737D16D8F2E0875586501C496B27AC970AE0741D3BB8AE4F39C14EA6A3D323E3BCF34D178E66EEFE4C202044A693AAED6B3F4963BFD34E7EBA350CF28EADD003981BC798A437AAFFC686C4A32B998801F0BFA749F9026686335CE51F8B33C882C713830A4`),
		Qx:    internal.HI(`AE4E36F0B48D62EF7F449FE517DF6331937EE58A027EC51E3C4BCDA2`),
		Qy:    internal.HI(`C2B77183CC4464DD857565687414C1E0330DB2E0438A2C74C411EF4B`),
		R:     internal.HI(`7C1F8F5C990D71166131533560D15702FC9AA9705A8EE315760E5EB6`),
		S:     internal.HI(`135436FF30A3D9A6C79668493FF6D1BDF314197E3C97E758C77EA023`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`ED7B09E3CD1A43853A0C1708BF3CA9D13C4201D32C3D9F86D4DC74838D23C98BD1E6D9432F8FA464CB36D1B49F880D8A1593D2192BF3F367D073E519889F20EE255A843096F72BD695E093B34555806C734C827B9EF28E1673DC852EE1741863669EF524D0259DCB07F9CDF911BFE802F8C742AAB5DFBAF7D21C866BDB159C67`),
		Qx:    internal.HI(`FC14126C4E3D7BCB84709F27E2FCF40D4704C4A67BC8C6C700F0C67D`),
		Qy:    internal.HI(`9398A5F379DC143CD3F50BC93E0A860CA888C1888D8EDA9BD2C92CD6`),
		R:     internal.HI(`1C751B2F77ED7A93397BD5EBD479C7CEFE76E67C424D338901520F68`),
		S:     internal.HI(`461A0B0D75C43F8D2AAE4D1F05A563F3D1DD8CBB42700906753CC029`),
	},
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-224)(SHA-224)_SVT.txt
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`CA3043DA50C6C27D5DB2E02A462500564FED6F6DA89A966D3A3C62BFCB19BEE37E4A61373AA0B45E01C9178AC0015742945694D220D06008EE705F70616F3AAC51991814C13F7D49CC92C4E146E706CC19B5B4B0EDC5AEC61D7612797F7BA7640F0006D5009B728433EF80A3B97179E93C2EA3CEFBF8499051F5C9335BDD6725`),
		Qx:    internal.HI(`F0DC7036BF5F01605EE00FCB4CD758A842B2BA9A0D8F04F764E8A2FA`),
		Qy:    internal.HI(`DE34AE18E20450C5D589ABAB62F51B80A1E80A968C3C9454426BE190`),
		R:     internal.HI(`A8EFAAFC438CB77A45DB4AE30084F106BD35A8873F0A132DD8B121DC`),
		S:     internal.HI(`D75556E3E368E46C2273D4EAD55AA1F29C4ECCC1340F5DAF0FBEC4E1`),
		Fail:  true,
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`352022FEC7BB412D917864B140C5661A39AA0529DD9A983969820A03C6DEADBF4D28142CC85CA688B8E101F8279EE178097DB2CC872C6D7D1F5EF374FC5D832E7D227D1538D5150E84E91BB8D2A0876BD144EF80BBE286E2B9B9DB7BC43CF77CE6EB001050BB1174176967121D769091F612C37136C609301E1F5B8BFB2063DC`),
		Qx:    internal.HI(`EEE40DB89FE8485F5B5CFBC442CF7C418F3829ABAE26572A5029B244`),
		Qy:    internal.HI(`CA641A1FBC5572A108FBF716B8747DF875FF9568AD1DD4AC6F4FFA30`),
		R:     internal.HI(`89D96098F55984CA756357F705D4655E292D6C9D15A1B9FE51E376A4`),
		S:     internal.HI(`87AB68719F6F3A70BC0498B2A9FDFE1F7EBB883D95FA1542163C0694`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`CC081C518967961AA361D0EBB977B16462FC73BFC86658BD86D2C38B9168A88F73C8E284A42979C13F506328E504BA8F16310022BF55728B4A6BE86101F063B7F73AB03C57B54AC64E3FE5C56145C1747AC13E356BD7BF0732166875CC2EB7EAD8227FD1D7AE5D6B490418C07A8E069741D3F6D47531A7376A8B58C7329A76F0`),
		Qx:    internal.HI(`51260B1180CAA77973D8A4B873117B907BB2B0868B0261B4B9D36F1F`),
		Qy:    internal.HI(`52A3CC6507EFD1B4EF4539B9ED7A474759B664650673446744AA9ACB`),
		R:     internal.HI(`E7B1076D0CD91A6322A38F7D2B9F0A0B5D2BB50FFE89BE6290232B4E`),
		S:     internal.HI(`5B6F23D25CF2E5BE01A97D4031E8F973CA3660291243940CDA92511A`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`0AEEA53F002F7A123A651B687EEF0386569DE66DF6BBAA864C765E5FF79D19B345442095DC6446F148DEE148FB0B4FCD004F58CBFBB47584A12849634F4C3DB3ED72B5C2C2C9A998593C271F3210950050CE1EEE627238CA9C7B9F39F3CEBE1CE70EDC9D7452052063E006DD99DD7A708D87371F92FE1824142480C47D978375`),
		Qx:    internal.HI(`F0BB84B705499383E8F6F2D081C5A1A041AF7F1C5B8D92AEFEA4027A`),
		Qy:    internal.HI(`7169161CCCC501A41503D829DA1D0F198A1598F7C020F377A3B3DDB5`),
		R:     internal.HI(`242C065D0EA6C30A42D505F1077055C5A7BBD462CB8676B23B3746FB`),
		S:     internal.HI(`58052327112BA6F0713920046F95B1BD0A242A9E86B7AD850568D9A8`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`5817E1FD1D440922D5A357EE782BDCC68DA78E3A88C4C7AB8114CDCF40FCAC3F38F031B36DDA3D1CCC284B38105B64D29B7C5353697250CA988805138D23194B3EF71B2FAEC00D9C1FCA5FC2FC47964F15CB6E1D21E0C5A1161609BA67767B02A0C51157777832B8CEC32225BAB462A408B2C71E5955922A51FB587FE200F931`),
		Qx:    internal.HI(`AB9A677E20F8FEF4B128481945E91B90D922E25EF4BA73474A5A1437`),
		Qy:    internal.HI(`84C7BEABEF7D20F7CB99339D9AB5C911CA240D74C6381CC511609950`),
		R:     internal.HI(`DDE575C97563CDF15EE17918A68F638133BC7CB63BEDB55277A181B2`),
		S:     internal.HI(`65271E1F368B3B74915417DDC67B5009B32AD008D70D58BECD47AC93`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`FD7156830DCC4D8457093CF0EE250125552E28E29F8D2447805E63374E2269A2487EA0B612660092D1722F6DC8A25BC9E35C51B314F21A9CF46DB0CA7BED06F85664F11E9FF03CF284A56B28A00372F890F0FAEF93675B9FADBEF82527B9291AE1384CEF59C33846CCCDA54059DCCD5C9972CBDC82F800450C7415FB794EBF86`),
		Qx:    internal.HI(`6BFC1EBAE132B5C752C30BF12514FF5BBBC5C787AE6337B607983A5C`),
		Qy:    internal.HI(`116E14D913F8B0C13CB42DF81C172178B11EF47245400012A1200D7E`),
		R:     internal.HI(`077805107478C66588F89B19467F8C38AA82F8E8D9C1759528C8A2D2`),
		S:     internal.HI(`8376A7E58D071779C93EE4EEA61DEA1EA163B7909AEEFF73F0502281`),
		Fail:  true,
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`6FE848EBC83D2E94C7423B97A6A5E0CAFECECCE9B73393CDFE90AF8D962165DF38939FBC60C835953F1FBE151E645BC7000F9B9CC6FDF09D9B47CAD6EF65C8BADA432D081164FB8A4FD3D49272F61DAF08EF776BAEB8088EEAA6A1C1E36753FDAB3B6AB47385610E6755DA34762211FDE054EAFCBBBF0E3E4FC927B47751C5A9`),
		Qx:    internal.HI(`C5F85C3DE288B53D0CCA6470D692B7C8AEEDFE998465CF25BBE81E26`),
		Qy:    internal.HI(`302930DB3045ED7E146F87F64556A7BBBA6D56C1B49B3A53FC2A2875`),
		R:     internal.HI(`C28D5BC80CF16B6894AECAA599C6F5CEDA7BB55BF0FDCA1D1B45BED1`),
		S:     internal.HI(`E2C7515DDC856B46439DE9ACCC7A0E99B0D27E1FD8FA63D7EF8C552D`),
	},
	/**
	// something is weird
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`09F0278869F0C4DCE157A56677CB38366FA57FA331E256141177A21B1686C6AD7F8556B914DA59EC00FA1CE44C87EDA3B77FA881F809846928D401A572A6680FBA15BAD753F97EBD7B9A54FBBF27B1A7ED7AB7DCC5B07F6A00A7D256E6DBA6001DB4551BBBCDF05E6C35911FB73922AB19BD72A10FAC0818D0839ED3BE85A25A`),
		Qx:    internal.HI(`F8C4023C6C7845B48BB9151ED6D229B426B216F9CFBF86C1E96987A4`),
		Qy:    internal.HI(`626A594ED756E77207EFAFCD843BAED9F91C202EB62F6879564281EF`),
		R:     internal.HI(`FC0F12466711078AD11E1E7DF40EA34BBCCBE199E2212091FED2D488`),
		S:     internal.HI(`0224E539B760F62D8BE1D7631758986DD0F36FC7F2F059C6E6079861`),
	},
	*/
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`A0B325A88271C42FF2A02AE4EFE4FD899A2624197CE55D1E23752D0DAD86E2533DE40040E7CECD6A1ED45FBB24A6D7D68AAB22EB45E60CB26AEFDAAC0182455616841B7077D9400B320C6B7C5D1C58DF2B5FC155D5E2DF9D622CB840DAD116E0E2B551EB536DCB47CAA95F2DC50388CD48278EF0A0B8773B3DD5103336B3D2BE`),
		Qx:    internal.HI(`6B3879306A318EC1C0E6EDF882CCD083AA270DFD4A7DE4F74A96A7BB`),
		Qy:    internal.HI(`CAC616F5F896ED87B2E5F8438C2222086F3FF0AE2B4EECDF55FADFE6`),
		R:     internal.HI(`5C11C271935D372A04299D337FEDF048D914BF72067FAE2B117FB49A`),
		S:     internal.HI(`84E240F93DED0C6AAE37C96F2945AAEE46C3D633FA7923548191CA57`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`20E22F456E039EAF9AB69E2F22580FD1737D1EC64F5E2B561A90B8BD71DC03EAFA7781BFC867D3CE52556518459F5EDE38ED070A98085931B450C9C05E810A57FA22EDD96D5DF63ECD85E1A1411EB8B0F5A19CDB88CDF676E2CE2934D2F47918655B495759325C3D90FA3999F57C66AFB1C7964E21F17CF96DF164BB2E513AAA`),
		Qx:    internal.HI(`801843669344A424D12446D046493EDE36474731FBD6F32E945925FF`),
		Qy:    internal.HI(`03CA33BEC75F1E95C02621BB3F88FC25BD1E3612D29BADEB8707A55D`),
		R:     internal.HI(`499AB6C23672EDE08642724B840AA45D712DD66EB24C58332F4B9DAA`),
		S:     internal.HI(`98846010E4567CA40F6C30D5527653937BF9418D9D477281650FCA38`),
	},
	{
		curve: p224,
		hash:  hashSHA256_224,
		M:     internal.HB(`0331A0901EA233A8CF5B71454AF76DF8F47522689CADB5A609CC9BDD9E82B12B640AC71B3DB0ADA88700451E0044A0E32583BB6DE192CB64CC7BC44FC5CBF9B22F00851096B162AC0EA3D976FE6B2B0FBFCF2648BA7E87095A55CEFD9BDB79FC7C093C89181D325AB58E984B8BD616A290341B180A095C7B3A05F869DCBD4BF4`),
		Qx:    internal.HI(`666937630E7AE8648B857A99775F01C8E6DF90CB9BA4BD077D9E5957`),
		Qy:    internal.HI(`EADB30154C7586668D896BCA3AB4D588B1813CC72A38CEFE7ABAEA2B`),
		R:     internal.HI(`22EDC71039068E4A69E209F03D50F30794C058F7837A964F825FD2FA`),
		S:     internal.HI(`4032741470C0F845EAB322839A40705513B33289E90BA18B9C560B04`),
		Fail:  true,
	},
}
