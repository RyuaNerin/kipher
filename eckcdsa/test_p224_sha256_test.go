package eckcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
)

func Test_Verify_P224_SHA256(t *testing.T) {
	testVerify(t, testCases_P224_SHA256, p224, hashSHA256)
}

var testCases_P224_SHA256 = []testCase{
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-224)(SHA-256)_SGT.txt
	{
		M:  internal.HB(`B39EEE40AA59E5CDE357FB3050CEF949185944C51F2183D0006C0F63499B14770B20DCDA65730C033C865A97C4D396F501796B88E777C8314295568A12EFDF477DC41588037C9B5C1573BD36C797BBB4AE5D1062AC0F9EDAB80428F65CC9FD6B0FBB438B0A23C67BC659B2A3CFF31FA2088DE486182B7EA21BCA95B44217E8B4`),
		Qx: internal.HI(`93175C7B0F9D0A10C760492A5A22B69E197EAE07BBFD80AE5A262AF0`),
		Qy: internal.HI(`B1E6B515DB0A9C0A100D478617C6792C13B0CBAACAA5ADC65F2A75F2`),
		R:  internal.HI(`88C16B91EA4265ACDC0C27DCA47AF42035D00B07A196D91D2ED41A67043B7961`),
		S:  internal.HI(`9A81C7738DE9E7D8FFBFE7BA97B0101E296EC2F3C52E71C908EF9017`),
	},
	{
		M:  internal.HB(`8DDE1A06483E4B6B6E1CCE609CE70144BE349E0D1D3669C8F09EB7B928ECA25A8BAF5609D2CD6DA0F249AED3379D27874B29FCF9F4C2E6F281DFB0DFBB85B6006758B98906F5F9F40311F94EEE64DEC1F0EAC7351F8A0531DF96A7ECED4836AE45AA10658A0542AB17515A085AB4FEAFB49CCE92687060FBE871D08BFCBFBDD5`),
		Qx: internal.HI(`BF08EA66D4E7B5BE99A8BDADB108565794D48601D42EC3DA98CD9AC6`),
		Qy: internal.HI(`EDA31C3AEAA44825C754B4DF4DDCA9533253E2617CF4817F3B603C7A`),
		R:  internal.HI(`B1F05010AED4CAFB46684190F94BB02777F3ED35FA99FA7567BEAA5DA7DEAF7A`),
		S:  internal.HI(`469F6EC475579860496DEDECF3A2AB3A78010A320953CD385DA530D0`),
	},
	{
		M:  internal.HB(`131F2D624D7426706C8D88215B30E475680CE7CAE17F162F3CC8625169A16D4E85D8EB2F83E109550AEF3FA1199E78A1A6B0F381E4EAD014479BA0B802C9F05A179F2E61AB11D7D6C04C467902F12728CFA6D93F120CBE593C91D227DB3C78A210E5E2778AF403D5231F6981D040E7660AB387765665991616FA4BE74E26C038`),
		Qx: internal.HI(`071EE0C81E67E5FED18D97CEBAE8F0B9B6B0EDCDF341C8AF564B3238`),
		Qy: internal.HI(`1BD626C1C5D88027C6C521C58A5601832783919F0A65CDF1949DED51`),
		R:  internal.HI(`2DB648E69439E8D7EC07F339938B6FDF0CE79F3113ACFFEE8CD3AE035D731544`),
		S:  internal.HI(`AC996076CCEEB5D8F2E3BBD88E4F58321A6AEFCF499BC5EC47BDB624`),
	},
	{
		M:  internal.HB(`7AC4F9D72E9F8C9F938C7D7882CB37224945F682E0A2A0C997CD632F00DD09942CFE6FCFED52F4E5385E61065FFC1D87457124A32B959C594AAA7B18DCDC2098C3FC4794657448C3FE08F8BDF863282B7FF51905F939DE1682D9FEA917C9558CA2CF8D467C921DBD9EA7321324B96D0B3F37E5B854AF3DB75A495BCC2C708222`),
		Qx: internal.HI(`06DF8AD2C1F6DFDBD8914E635B2E34650B1D32E950E05EC01C719C8D`),
		Qy: internal.HI(`4DD2D2C7D24B5B000EB1034DF11F28A26EA1AB488FAB9BA426BDE60F`),
		R:  internal.HI(`527DCA791D017A41891C1CF2A654F993F8C916BB078071499C60A12C4E0EEE7A`),
		S:  internal.HI(`487B34EFFF81A2E066D272FED7CE83FDC172DA53A22A41EFA9DC4C95`),
	},
	{
		M:  internal.HB(`F73254EA5F638FBC97FE026803DD8C8D96439DB88D431A5AB4910D57E0C50C72B786B66D84C641152E786606FDDAAA7C5DD063E53E665E873CF295043DE5DAFE9DD4D8A7A9C162817329641BC2DE760F353B5D0B49B5782C64537D78971461AF31CDE555D583A4263DCD09C14A4426E2878ABADBD8F161A26843543F8CC09AD5`),
		Qx: internal.HI(`3E776F0E9815514CA8592456525981F156DA8E9C08994E22C1F747B6`),
		Qy: internal.HI(`FE2EC42253D099DDA16ED33C0678CE50B14DABE7EAC1BE452D26CA9A`),
		R:  internal.HI(`CF9E535C047C24F8FC7B88F46086BBFF9B9239BA2235D833DFE9BBBC72573F29`),
		S:  internal.HI(`709D4EA896C9A752CC91209BBB3DC1D70B9323C6D8F9960C9651636F`),
	},
	{
		M:  internal.HB(`BDCD101E5363448C2BC76AF7D48B79FB836AB1EF5C0499A649F9B6CDFD7C072B57D3958EBCDF06A8A223A5A4E85CB4C6223183CA90022960D356437E1805B4D1DB8AB61DEB9D37D4D393DE9A5687A61825DC76D57524A06098E2A4964F40314FEF43BF29086DABD4B37543903605A52F1512DA6255CF189BF4CB8B425F3A9A98`),
		Qx: internal.HI(`42C04E403F7BBEB478F94F764DFEB516D59A0FB4D21DDC7BDDBCEFCE`),
		Qy: internal.HI(`FB6E644F8A0C340E1B53178F08BB401EAE7E61F5DC14DE44677E8829`),
		R:  internal.HI(`0806146AC8DFE3057B6D301C2AC499ED0D09277EF0CD42AFD8530DDF6BE2004F`),
		S:  internal.HI(`E1896FD0036C5F0FFC6C8CFEFD532397F1BCF4BB3ADE93146F9964E3`),
	},
	{
		M:  internal.HB(`00F803F78044BED303CB0A26E7F790AF431F06ACC28B307108E8B1954C269103414ADEB60842556348416FE512A5CDA7C7F859D5950B11AAC1BBD78C61623F53B083B37B9FACDC80D12AB93CA6804A8A833D3BE7F12B6B74D06AC709327259B11294EE458AF2478CB5823384DB207E351E311BD33FED7666B1C552DB9C0316AC`),
		Qx: internal.HI(`9848F17BB5F119E3B72AE045971DB49532ACF2A06222EEBBC73F71C3`),
		Qy: internal.HI(`6A35C75460CE486EEA5B912C510C54B321FAB9FDE0722ABFCC82C369`),
		R:  internal.HI(`97AB6BE0F3AB4B10AEC73596ADDDB658B69296C04CB906CFE28277A9297F9CAD`),
		S:  internal.HI(`6884D3B984D052EC2B664152D0140A4814EF7BE2DA906D5299F5B76F`),
	},
	{
		M:  internal.HB(`F41801FA58AA1354D3ED35FC324667EF0BC46F74317BF37FA74351B4C1E93B3FA94E6768DD924309D2B71ACD70DB8A64818AB98CC2272A27BB04A6300D2012CA5022A5443992654821D24906A8EEF8A882C07EC4316DEB2DC1D03AD535CE6C17CD25462ECFB78C12F6DA2EA02FB9453AD64D4FB00BEF8FC75515FE0D343EA457`),
		Qx: internal.HI(`5FED28D5BCF32975133B287159DA7B5ECBD889D1D914C5DA3C310D8B`),
		Qy: internal.HI(`6DA48A1F7C9F3FF493A5A4EC3FCE49471408D3E5A6025A2CCC1F9046`),
		R:  internal.HI(`C7ED3E05DF2D578D0C0EB3D194CEC1A69C85BF8F4CD303EEDAA5FDD7BD51B214`),
		S:  internal.HI(`E6A2348F7266F290B70B11908F532DC289773A39799C1C530B12ECBB`),
	},
	{
		M:  internal.HB(`CD90DCAA503855D350133F7BA79C90FD0EBEC2C91F78F695D8EDEB2C4FE69C22C4450328AF75E45FF768FA5FF72080408449767189F9899D751504701062BF7AF0CD5FFD2CF2E6F0786EE3FC4EF542B757CB14F2AA8E364F20F751FC4B77FEC7545A9B674B608E292A5F88E925F48E8070C94A7E2C7A788292A0E3DC1E10D5DC`),
		Qx: internal.HI(`BFC3B9E5A5EA336EFAE65ACF4D5BF46251B92658CB7C7716A8B10AA7`),
		Qy: internal.HI(`EADFADFF677AD028588A4F2CFB211E0E0F449834ED8055C36055B812`),
		R:  internal.HI(`3CCBF932688CCDBB34E2F69FD2E49074BC1E34F86C77BAB00BE970FE7E310034`),
		S:  internal.HI(`CDB3802240DA809E6B53B4E3DA524F5CA74197012C121A3A1DC151B2`),
	},
	{
		M:  internal.HB(`BFC56A8CDB9499162E1F7DA83B1EA11E8172D03000264E7550CAD403EA4346F0C590877AF28D4C28683A62A0999A4281039A6509602541CFA3D3454F5D4DDAA5C2E5B528ED71723D89E3DA218EBABEFB35C0D1F4CF33609F9FC35F836991A505DC96C1747290619606F69363B1F5ED4B2109E2C21631445B1E4A554C4B9C407F`),
		Qx: internal.HI(`C0798EC1DF90ACA138D9E14ABE11B8E8F19EFAD36D01DCF2B7DAFF9C`),
		Qy: internal.HI(`2C2EFD00D94F1FF5A7C26E3FD3328EAE042DC91F05C6D1088E476575`),
		R:  internal.HI(`DC6069B3360962D085F4EDDBDA37E530673E80802578C53735B5E45655530985`),
		S:  internal.HI(`18475B2A0E55D692B8E60D9BC72BBFA6B7406A903B346061541C7C9C`),
	},
	////////////////////////////////////////////////////////////////////////////////////////////////////
	// 암호알고리즘 검증기준 V3.0
	// 테스트 벡터
	// EC-KCDSA_(P-224)(SHA-256)_SVT.txt
	{
		M:  internal.HB(`66305751A2D2C7084DA6E499D498D1577F8A5542DAA4F154F163EA19A31BD690203DEFE7D12ABC16B1D344DD68FD7C33275AFDA9975AC8D2CAE62A26A7D879809A77BFAA0BBA0B0C73A191FE86662BCA47D383E77F7F716F7A42C85D6A9E32986371455755AC78E5E55C1C4298D713C12E3E4C0D9F32D3BB94FB78DAD8E6D890`),
		Qx: internal.HI(`C22ABA434CEB4418DE43FFE715AA49260FD4C3FECFCA9E6B1ED6F30B`),
		Qy: internal.HI(`1425D8B6C7695B6062242BDBFD2BD2E4ADC2CFF34A62D27516EEBD9F`),
		R:  internal.HI(`65353BB7BDFB3A59D403EF1A59D6936E93F759D7EF17CDE6FAA088A47DD9DDC5`),
		S:  internal.HI(`C65C240F8F6115CD612A40727F76E83E63B0713BB201E050D2E0A39B`),
	},
	{
		M:  internal.HB(`2ACC54BF120AC2DFBE029A0F3EAFE3B94D85D30A791B6D35FCBB61E6108B49C10664AD61E616EFC0E4249C656DC3254919239DDA43C0620199D75E9464B3CC1BD2E7195A0A33DE72552D881B215BA5C922E27335647EA1258BDBE4E411004A4DC3F4F5D148442A65029146BAFECAA6AA524707449C5D5A54F79DAEE631691A8D`),
		Qx: internal.HI(`A2B44154F1C857FFF14BA2144019D81111C75B8C71EA1CA1A30EFC2C`),
		Qy: internal.HI(`A3F8DC06BC1CC8B0FEBC1A2BDB05BC2C72352FAA817F59E78E11C498`),
		R:  internal.HI(`E82127B8DA7E054B392E3451034FF811E85F5694673DF00310B1BD0BCC827EC9`),
		S:  internal.HI(`09121E872738C5D1DD4CE4BC30B021E7F72D125A8938453A81A0E7DF`),
	},
	{
		M:  internal.HB(`F447F553739590408A65D391AD96E8FA9FB23B90A9E2CC68BD20ADC32EBD99F119BCFEC4320A74CF1A1D004CE70704D3007655CEF8B84FCFA14658AED8DDB8A05BAA2B1D3A9B1AD60E4EE38CA5B90AFE44FDD7CD29822C8978D30CC43649B8597C0EC3C88D8212E4F77F2396598E8A0412980AEA611035593A23F6BC60FAEF58`),
		Qx: internal.HI(`39618DE3EE40DA4DEB991A7947DF890C5FD24853AC43D84C68EAF6D2`),
		Qy: internal.HI(`5AA123DDCE69CB67168EF96769460452430EAAF496D3BE021E4306E5`),
		R:  internal.HI(`452B9C243C5F4C867EB3EF081E0A23DAAF8E4959B97AA6B1876DA389445B17B8`),
		S:  internal.HI(`834C945368AFB414E79C3C0D901D45E582C31C569A13B538A16454E9`),
	},
	{
		M:    internal.HB(`8F1619752474F030DF7494C093A0D6D641426B0E58C869A33A4B37B5103C4341787B478409E15ADC468D9CDDFDEDCDC699DC80E677C84D3DC24ABCDED5052751C1517794E580E7E63431AAFD299DAADDEA393DE6B6A200541036BBC9B93878110600B1C69202C801D129A1B39A6197E125C5A055824BE5011F5D8DB0544C142C`),
		Qx:   internal.HI(`341A990FF8CCCB76622A63F4BB068D30DF0DC4E1051201588724DF18`),
		Qy:   internal.HI(`F249CD4FF540911DB678FF836D5FDD006E7C84EA1B7DD110A6C8E233`),
		R:    internal.HI(`CBB7784EA4DD67442FF3155C7C6CE1E4F4C5DC3EFA69B93AE6FD07426A1B5A9C`),
		S:    internal.HI(`427E612C373464F490B4CE0633A6743DEEBFDF335D08DD6DE8132B0C`),
		Fail: true,
	},
	{
		M:  internal.HB(`6C43150627EBA56D9CBFF3FDF9286B01248330D64241ED0D5E743BE3332CBD8868509C48335C8F1ECDA40EB2988EE14B5A66AFC48A30B805012E580E006B8D11EF0F12B4D3C94EF2B30FAAAE0B83847B8E4365A62F679700DD83D87CC7603524AAC9D967F9A2D45158249AB446705BC26C56D8296588E3377E0B300543299E45`),
		Qx: internal.HI(`6EC218A9EBCB26CF2EB30EA4AB11462A6934A1EEB88AF7295D6EB034`),
		Qy: internal.HI(`0F804E2914AFDAB6A7CB8AF4A1D0B8A8CBDD3FFD8E05EBFB54EECBEC`),
		R:  internal.HI(`5ABBA3076EAFB1AACF9A4A5A3AF8DE7D1B845CD42C6029EFC0F0E68C3B1164C3`),
		S:  internal.HI(`15FAE9C5B0CB355D4D4C0CB8413B61BD1FAB1D0F3E45FC824F29DC3C`),
	},
	{
		M:  internal.HB(`B09BA11EBEDC4ACCEFFBB1C22632F9A6C7D526E89EB88F7B4622ABB04A78A37FDE4ACF405BEA95286431BA81DE9DA3FAA37399FC1D0095709BB696D4B03C130F551E81E92246928F5FE60B0E38FBE829E10FAA6DDF667A035D3A1700B12C11C0009B742821054729A40F097C3F6E9AFE5E4FA79128641CE1C147A5050A4FC560`),
		Qx: internal.HI(`AB4D447C6C569A4E4244382BB0D2ED028FF1AB774027859F49EA2389`),
		Qy: internal.HI(`C637D7084E85CEA40A991621F900378C406360E00171B22B32618FE7`),
		R:  internal.HI(`68129233E06FE26AA3BB35108F783B3210089349C18472345F3A0B5419BDB664`),
		S:  internal.HI(`209D9DEAC7F3CEDFEBBEF288E9C098A9CDA1DBA898F945430A5456BA`),
	},
	{
		M:    internal.HB(`8A8997CE026528DFE8406463E5334FE0210643A238FCBFE8BDE247DF77FCC721FDCB560ADE44C5A5F566E87D14D4F517BB67E861E8485BDBC397A9BE3C2AD4F34E6D0C1B06E241710155B4CE7847D3969B906BE044D27A8ACAB39B9C006E6FB01EC9EAB0ABE7E6A82014F90056FDD7DE866F4202E217E9CC884F12D5EA8D87CB`),
		Qx:   internal.HI(`139C5F6B8FE932B74D65241383E42F9FB7AF1AFF05C12B26328B35B0`),
		Qy:   internal.HI(`43EE4AE30188815F1BA244F0FF147BCD9967DF18F606B61BE8D1AC86`),
		R:    internal.HI(`DFE04785116A4ADB75C00AEF2D63A2E6BFFD104E989AF90EFCDF4CDBA4425335`),
		S:    internal.HI(`FD836D9261D90FB52DF9857BA0B4548E7168C5A7DD28A0C413F963D5`),
		Fail: true,
	},
	{
		M:  internal.HB(`CD65B9E6E7B8143F7E81A7E4F762A4820A6CAE25270DB3DB2D177E0FF439DA47B4EA8469D615BE31377FB800534CB1AB8281016FAAF6E3BE8FA440FE2BA043E62EC9FDAD227BF0A13E37A9596DAA582896F78BFAA9ADEF407E534A1104EF8833D7731326B2754CE3DA1B6371F7A657007483719EDE3EEF58515A38B285565B57`),
		Qx: internal.HI(`C00A3CEFCAB738A9234D8518AE5CB832A9804A4BE54F5DF651892BD0`),
		Qy: internal.HI(`373BA460EE96B0B20201C3C29A9F256B01E0D2A27175550D2522092F`),
		R:  internal.HI(`BCA234ED76FD37AE80FF36EFB4EB3B74E39EDFDAAE82A652375B02037B87BCCD`),
		S:  internal.HI(`5606A7B5EB6B0C154769CFB0172D5720A4BCA99884F50D0D14AD5376`),
	},
	{
		M:    internal.HB(`E01C0D15935F92752536974A75557D75613EFAADC0E347897710115C0B57C5E891D867836F9E92A5D7DDDE6951DF391555E474214FB393D0F13D27DD4FFD374C8C6EFDDBAB640F263C7724FC7E6C11FCACBE3BEF102F0EF59797B60667B7AED9646D62A6EEF47EDDE78F0D55465929F93A9046723210CFF1BD58457EA5D873F2`),
		Qx:   internal.HI(`EFAAC8B9A59DDC0D7D68B008446EAAC1105A1B0DE4BB6AEAE3750D84`),
		Qy:   internal.HI(`FAE9104E586E04A6D40B08C5A09606F75A3306BD1EF333A7BE09B213`),
		R:    internal.HI(`910926A8F20D8A217A35E96F5AA2974C968027154A27A7FD6F6CC54681D727CB`),
		S:    internal.HI(`C3C64DCC32644A8A953465F2E099DF2E0BABC0C7DD6CC95D8C717A2C`),
		Fail: true,
	},
	{
		M:  internal.HB(`8F49C1156F48C57A8468C7D1ED50517045E50CDFEC9B620E0B83F70627B903FB8157162F189C3F83886B363094C97E73A4860FB900B28C11021196612A27373E2517F2D2C9EB8EB3287306198E6CC65BB00BA710DF96B9D16E3D6BEB32C807004A5E1DE806EAD7DA2E65C919799F383C6DAA7052839F22E5CAD05A35D1281897`),
		Qx: internal.HI(`312E9F2CF8EA08319AA670808B43C8CD664A54667A396523827DF2DD`),
		Qy: internal.HI(`6C0A226FC334F58E1492BDA8867EAAAE4575CD4E80B200FB477E5E5D`),
		R:  internal.HI(`5E6FCA7C5BBBEEC4996F9C457F0E162819364A8EC5BB32A0C3F8AF655030F663`),
		S:  internal.HI(`83BF870BAABB26B22AEFC6942FCE657BAB08FAF5A98EC54490A73C7C`),
	},
	{
		M:  internal.HB(`C606633A946753CBEEB8C3DB0F0AB34463AEF3CD6B276BFD5CF2D6811633E65EB1C43BBA2DB3009782B4B222B29BE409FB4D7283E4500A5338666CA5AE4B55F24B1872C341564C935F4989A4C7717BF803E05456384FF08187969F85507BD35485691177A317292A89A6F49FD23BAEB3D0A225DEE0FCD2A42F071F55A805D85B`),
		Qx: internal.HI(`8DE419F60C0957ADF0FC639DA05F55AE9AC1B45083AD978C76CADC09`),
		Qy: internal.HI(`4AE23F7CEEF01C13F171DDCD2CA73607D641D43CCE491032E437944C`),
		R:  internal.HI(`BD8A7B2E3EFF25BC6776D7626C791FA5CC1FAC056B553B84ACF5A774D08E842A`),
		S:  internal.HI(`B2740E187A866EDC41F6CFC9FBDB671EF33905B25778BC8533450A77`),
	},
}
