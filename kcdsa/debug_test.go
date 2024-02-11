package kcdsa

import (
	"testing"

	"github.com/RyuaNerin/go-krypto/internal"
	kcdsainternal "github.com/RyuaNerin/go-krypto/internal/kcdsa"
)

func TestDeug(t *testing.T) {
	domain, _ := kcdsainternal.GetDomain(int(L2048N256SHA256))

	kcdsainternal.Verify(
		internal.HI(`8C731D29D00471C0F8DA463D9B6E5CA6665C4447B189FE9032CFB1274A0277C373574AB434B589799438A383101543FFBD0FD41F36285D7718E7E00F22BF623DAE474FD4B0B395BCE51A205FAC6893DE66BFB401BE09433DFF253240774249CDBEF6173D53583DAF50E647C86786FCB25B1DA0B3F170C0BBF5FD193C86DD17B9D8FAB6DE38E808B4EC2E1E3C861C7A71494942116F66E42873EED4C7AF49D36F2F1988196BDD5800F79680744301017E98FE2C609194A3C226048B0D29E047585617549B73E7A79A7EA2E6D8234D83B4B46D58AB6CE10842EF914C0A03F40C2C53AA33BED6E1594E96AED1AA44471877DAAB1B1FABBA156FCDF63EBF58088267`),
		internal.HI(`854A9FFDBA9FC911FCED53B1600EF60DC5775431D0A5635F4B85119675E31673`),
		internal.HI(`7DF978D12A38217D7187FFEB99DF28A1ECE438F0A52242CB648628BB795DA653B3597300D475DEF62FA06F3E7A849CEB516006EA985FD0AE837888B7C7C7471181DD939286E5097C333423FBFF3DA3DCF0714E1307E1F37E163241B3786CA1B4BD74BB48B6E6E8C888E31CE440BBB42D7E37589E154E9B2E02A4A7D53E6A0DAA28F15ECEFC36559D54E624520DFCB7F8121FE0635565836C68F66CCFC9AD071A653D78AC5C4E890FB0821EB2668B2FD2A75EE0D1BD5AF80E0A0769EDA3E248E1F090FE668E0685A4CCDFBFBEA67E7DCCB0BB494F1E85AEBE63548A8E96D83B56B626E38EBF6E9310A1BA90731BDD701FF0416C16B8B2B3FF4E98F7C3E96C906C`),
		internal.HI(`2B250EC6B871326103EA2BE09F862E51F8D08638F089EAA74829835E1232A649231E5F7D11C2EA0F97FB01613F3B51B812C51B92B066014A8A6A863DA1AFCB932638D28EBF21FCF611942A4A2C3F2AB30E820A83FEF9A143CF118D44C60EEABB5F1EB41510AF8DCB651053533565A8D809A61D65A26BDA7E0274628995D8ADA707F1C6983935E65928430DC2CCC0AE252CBDE954B75F6C1E91AE656A9D39E707E26386500F6E4F7985A56161E4C2A7F3B9341314538E92CA736B77A9341962BED8FE75D2E6234C5BEC41CCD821ADE5720C916D3A829975EF59BFD88689BDCF8EAF0A0E5408E5EEEEA8DC4DCFC4EA4A3228E58728C18D2A8E33434A4823958461`),
		domain,
		internal.HB(`E015626BE12AAF67BCC0AE67E47823CABEFBECD368E007C6E0416C6B2B689E8F7B8C23A7D5EFBF484B56AE9C4B3DA96F0CAFFE0AA22B817AC07FA178AFF767858AE1A97E36AA342D34A13928BFA9CC804FD9A97F02611E16F1AC5226FD15B9431A27F727A1C8056FB238CED4264D1E872AD225CDF9744C95C224B32350541A77`),
		internal.HI(`0037279FA2C1E545976C10356E88219DF3CCC477761114C82846A46F922182E3`),
		internal.HI(`7152AB648E2E3616CF2002BD8CC9984B99CF3365787B97F8E691A618057E0FF3`),
	)
}
