package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	//s := "{\"id\":1,\"message_tag\":\"\",\"message_tag_title\":\"\"}"
	//a := strings.Replace(s, "[", "", 0)
	//a1 := strings.Replace(a, "]", "", len(a))
	//fmt.Println("s la ", a1)
	//{"message_tag":"NONE","message_tag_title":""}
	s1 := "[{\"id\":1,\"title\":\"Giá được giảm còn: 310.000đ\",\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/XJZ1hs1gxR_1517816830.jpg\",\"url\":\"\",\"descriptions\":\"Giày cao cổ nhung Hàn Quốc (đen)\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt ngay\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8127\",\"script_name\":\"cao cổ nhung đen\"}],\"tags\":[],\"payload\":\"key1519617729ctPhJkguQMC\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8127\",\"script_name\":\"cao cổ nhung đen\"}],\"tags\":[],\"payload\":\"key1519617729kXKn43BP3St\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}]},{\"id\":2,\"title\":\"Giá được giảm còn: 250.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh Hi Hàn Quốc (trắng)\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt  hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8133\",\"script_name\":\"hi trắng\"}],\"tags\":[],\"payload\":\"key1519617729qNK0ILbJHnd\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8133\",\"script_name\":\"hi trắng\"}],\"tags\":[],\"payload\":\"key1519617729xEWSVrIH33F\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/gK9ws9tZLF_1517829460.jpg\"},{\"id\":3,\"title\":\"Giá được giảm còn: 250.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh Hi Hàn Quốc (đỏ)\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8146\",\"script_name\":\"Hi đỏ\"}],\"tags\":[],\"payload\":\"key1519617729OqJR2zUU2lI\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8146\",\"script_name\":\"Hi đỏ\"}],\"tags\":[],\"payload\":\"key1519617729NlqzwahUoPz\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/OmTdR7PE5Y_1517829559.jpg\"},{\"id\":4,\"title\":\"Giá được giảm còn: 260.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh cá heo Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8135\",\"script_name\":\"Cá heo\"}],\"tags\":[],\"payload\":\"key1519617729kHuMqcQtREU\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8135\",\"script_name\":\"Cá heo\"}],\"tags\":[],\"payload\":\"key1519617729BwWbrRdMz18\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/txCe8t3BIw_1517829700.jpg\"},{\"id\":5,\"title\":\"Giá được giảm còn: 260.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh xương cá Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8136\",\"script_name\":\"Xương cá\"}],\"tags\":[],\"payload\":\"key15196177290BFbSzzH9Up\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8136\",\"script_name\":\"Xương cá\"}],\"tags\":[],\"payload\":\"key1519617729DGPQxjIcejI\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/X01QPsdYJe_1517829816.jpg\"},{\"id\":6,\"title\":\"Giá được giảm còn: 260.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh mèo hight Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8137\",\"script_name\":\"Mèo hight\"}],\"tags\":[],\"payload\":\"key1519617729au92IWCK5Cl\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8137\",\"script_name\":\"Mèo hight\"}],\"tags\":[],\"payload\":\"key1519617729KNeknWthl7Y\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/yPl5TMKyDR_1517829902.jpg\"},{\"id\":7,\"title\":\"Giá được giảm còn: 260.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh thêu hoa hồng Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8138\",\"script_name\":\"hoa hồng vàng\"}],\"tags\":[],\"payload\":\"key1519617729aYwuGIIZqTu\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8138\",\"script_name\":\"hoa hồng vàng\"}],\"tags\":[],\"payload\":\"key1519617729AVdxxYCaknY\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/6z7Qvew3G6_1517829989.jpg\"},{\"id\":8,\"title\":\"Giá được giảm còn: 260.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh thêu hoa hồng Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8139\",\"script_name\":\"hoa hồng bạc\"}],\"tags\":[],\"payload\":\"key1519617729yIlucDQkC1i\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8139\",\"script_name\":\"hoa hồng bạc\"}],\"tags\":[],\"payload\":\"key15196177298vYQdYh6tS2\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/iDtytW65be_1517830080.jpg\"},{\"id\":9,\"title\":\"Giá được giảm còn 330.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh NEO Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8140\",\"script_name\":\"Neo xám\"}],\"tags\":[],\"payload\":\"key1519617729GpAF2KZpJx8\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8140\",\"script_name\":\"Neo xám\"}],\"tags\":[],\"payload\":\"key151961772952kISEkTWtp\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/pFSWHu2PvH_1517830210.jpg\"},{\"id\":10,\"title\":\"Giá được giảm còn 330.000đ\",\"content\":\"\",\"descriptions\":\"Giày xinh NEO Hàn Quốc\",\"buttonList\":[{\"id\":1,\"title\":\"Đặt hàng\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8141\",\"script_name\":\"neo hồng\"}],\"tags\":[],\"payload\":\"key1519617729VkCTBJmtM3T\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Phóng to ảnh\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8141\",\"script_name\":\"neo hồng\"}],\"tags\":[],\"payload\":\"key1519617729mIV7IlQftVi\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isShowScriptSelect\":false}],\"isCurrent\":true,\"image\":\"http://static10.zamba.vn/static/rdchat/media/2018/02/05/esb7dcahHN_1517830350.jpg\"}]"
	//s2 := "{\"title\":\"{{fullname}} có thể chọn mục \\\"Sản phẩm & giá bán\\\", Album ảnh giày trên fanpage, hoặc vào website để xem nhé, rất nhiều mẫu giày xinh chắc chắn sẽ làm {{fullname}} ưng ý đó.\",\"titleFields\":[{\"match\":\"{{fullname}}\",\"origin\":\"fullname\"}],\"buttonList\":[{\"id\":1,\"title\":\"Sản phẩm & giá bán\",\"type\":\"script\",\"content\":\"\",\"kichban\":[{\"script_id\":\"8172\",\"script_name\":\"Sản phẩm & giá bán\"}],\"tags\":[],\"payload\":\"key1519617729lw4XZrBZZw4\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isHover\":false,\"isShowScriptSelect\":false},{\"id\":2,\"title\":\"Album fanpage\",\"type\":\"link\",\"content\":\"https://www.facebook.com/pg/canmuagiay/photos/?ref=page_internal\",\"kichban\":[],\"tags\":[],\"payload\":\"key1519617729GW5BcOKBZa9\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isHover\":false},{\"id\":3,\"title\":\"Website\",\"type\":\"text\",\"content\":\"Đag cập nhật...\",\"kichban\":[],\"tags\":[],\"payload\":\"key1519617729y4ou4FhegCL\",\"app_facebook_url\":\"\",\"app_fb_attributes\":[\"\"],\"isHover\":false}]}"
	//s3 := "%7B%22content%22%3A%22Ch%C3%A0o%20%7B%7Bfullname%7D%7D%20m%C3%A3%20code%20Haki%20T%E1%BB%91i%20Th%C6%B0%E1%BB%A3ng%20c%E1%BB%A7a%20b%E1%BA%A1n%20l%C3%A0%20%7B%7Bgiftcode%7D%7D%22%2C%22out_of_quantity_message%22%3A%22Hi%E1%BB%87n%20%C4%91%C3%A3%20h%E1%BA%BFt%20code%2C%20b%E1%BA%A1n%20vui%20l%C3%B2ng%20ch%E1%BB%9D%20%C4%91%E1%BB%A3t%20sau%20nh%C3%A9%20%3A%29%22%2C%22event%22%3A%7B%22id%22%3A%225b0f9db309f069118b73a533%22%2C%22name%22%3A%22haki315%22%7D%2C%22giftcode_type%22%3A%22event%22%2C%22prev_data%22%3A%5B%5D%2C%22giftcode_api_url%22%3A%22%22%2C%22config%22%3A%22%22%7D"
	//s4 := "%5B%7B%22id%22%3A1%2C%22title%22%3A%22piao%20liang%22%2C%22image%22%3A%22https%3A%2F%2Fkenh14cdn.com%2Fzoom%2F600_315%2F2018%2F8%2F2%2Fphoto1533181234200-15331812342001061006010.jpg%22%2C%22url%22%3A%22http%3A%2F%2Fkenh14.vn%2Fkho-tin-voi-hinh-anh-duong-mich-gia-chat-xuong-sac-khong-phanh-ben-canh-dong-nghiep-hon-tuoi-20180802105055665.chn%22%2C%22descriptions%22%3A%22Kh%C3%B4ng%20c%C3%B2n%20l%C3%A0%20m%E1%BB%B9%20nh%C3%A2n%20g%C3%A2y%20s%E1%BB%91t%20v%E1%BB%9Bi%20h%C3%ACnh%20%E1%BA%A3nh%20tr%E1%BA%BB%20trung%2C%20D%C6%B0%C6%A1ng%20M%E1%BB%8Bch%20gi%E1%BB%9D%20%C4%91%C3%A2y%20xu%E1%BB%91ng...%22%2C%22buttonList%22%3A%5B%7B%22id%22%3A1%2C%22title%22%3A%22Script%22%2C%22type%22%3A%22script%22%2C%22content%22%3A%22%22%2C%22kichban%22%3A%5B%7B%22script_id%22%3A%2211960%22%2C%22script_name%22%3A%22KB1%22%7D%5D%2C%22tags%22%3A%5B%5D%2C%22payload%22%3A%22key5BiOhaqi1533112059646%22%2C%22app_facebook_url%22%3A%22%22%2C%22app_fb_attributes%22%3A%5B%22%22%5D%2C%22isShowScriptSelect%22%3Afalse%7D%2C%7B%22id%22%3A2%2C%22title%22%3A%22text%22%2C%22type%22%3A%22text%22%2C%22content%22%3A%22text%22%2C%22kichban%22%3A%5B%5D%2C%22tags%22%3A%5B%5D%2C%22payload%22%3A%22keyXJLCZvzh1533112071846%22%2C%22app_facebook_url%22%3A%22%22%2C%22app_fb_attributes%22%3A%5B%22%22%5D%7D%2C%7B%22id%22%3A3%2C%22title%22%3A%22link%22%2C%22type%22%3A%22link%22%2C%22content%22%3A%22https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DGIDoQsQyS0s%26list%3DRDtFuWRRne-Oo%26index%3D11%22%2C%22kichban%22%3A%5B%5D%2C%22tags%22%3A%5B%5D%2C%22payload%22%3A%22keyIo5PXZv41533112081685%22%2C%22app_facebook_url%22%3A%22%22%2C%22app_fb_attributes%22%3A%5B%22%22%5D%7D%5D%7D%2C%7B%22id%22%3A2%2C%22title%22%3A%22Nh%E1%BB%AFng%20phi%C3%AAn%20b%E1%BA%A3n%20b%C3%BAn%20%C4%91%E1%BB%99c%20%C4%91%C3%A1o%20%E1%BB%9F%20H%C3%A0%20N%E1%BB%99i%20m%C3%A0%20m%E1%BB%9Bi%20nghe%20t%C3%AAn%20%C4%91%C3%A3%20th%E1%BA%A5y%20hay%20ho%20mu%E1%BB%91n%20th%E1%BB%AD%22%2C%22content%22%3A%22%22%2C%22descriptions%22%3A%22Trong%20nh%E1%BB%AFng%20ng%C3%A0y%20th%E1%BB%9Di%20ti%E1%BA%BFt%20m%C3%A1t%20m%E1%BA%BB%20th%E1%BA%BF%20n%C3%A0y%2C%20th%C6%B0%E1%BB%9Fng%20th%E1%BB%A9c%20nh%E1%BB%AFng%20phi%C3%AAn%20b%E1%BA%A3n%20b%C3%BAn%20%C4%91%E1%BB%99...%22%2C%22buttonList%22%3A%5B%7B%22id%22%3A1%2C%22title%22%3A%22Phone%22%2C%22type%22%3A%22phone%22%2C%22content%22%3A%2201697653197%22%2C%22kichban%22%3A%5B%5D%2C%22tags%22%3A%5B%5D%2C%22payload%22%3A%22keyEmfIdEfM1533112180956%22%2C%22app_facebook_url%22%3A%22%22%2C%22app_fb_attributes%22%3A%5B%22%22%5D%7D%5D%2C%22isCurrent%22%3Atrue%2C%22url%22%3A%22http%3A%2F%2Fkenh14.vn%2Fnhung-phien-ban-bun-doc-dao-o-ha-noi-ma-moi-nghe-ten-da-thay-hay-ho-muon-thu-20180731232155452.chn%22%2C%22image%22%3A%22https%3A%2F%2Fkenh14cdn.com%2Fzoom%2F600_315%2F2018%2F7%2F31%2Fphoto1533053987039-15330539870401300523722.jpg%22%7D%2C%7B%22id%22%3A3%2C%22title%22%3A%22M%C3%AA%20%C4%83n%20th%E1%BB%8Bt%20g%C3%A0%20nh%E1%BA%A5t%20%C4%91%E1%BB%8Bnh%20%C4%91%E1%BB%ABng%20b%E1%BB%8F%20qua%20con%20ph%E1%BB%91%20chuy%C3%AAn%20b%C3%A1n%20%C4%91%E1%BB%A7%20m%C3%B3n%20t%E1%BB%AB%20g%C3%A0%20v%E1%BB%81%20%C4%91%C3%AAm%20%E1%BB%9F%20...%22%2C%22content%22%3A%22%22%2C%22descriptions%22%3A%22V%E1%BB%9Bi%20nh%E1%BB%AFng%20t%C3%ADn%20%C4%91%E1%BB%93%20c%E1%BB%A7a%20c%C3%A1c%20m%C3%B3n%20g%C3%A0%2C%20ph%E1%BB%91%20H%C3%A0ng%20H%C3%B2m%20chuy%C3%AAn%20b%C3%A1n%20%C4%91%E1%BB%A7%20m%C3%B3n%20t%E1%BB%AB%20g%C3%A0%20l%C3%A0%20n%C6%A1i%20...%22%2C%22buttonList%22%3A%5B%7B%22id%22%3A1%2C%22title%22%3A%22Share%22%2C%22type%22%3A%22share%22%2C%22content%22%3A%22http%3A%2F%2Fkenh14.vn%2Fme-an-thit-ga-nhat-dinh-dung-bo-qua-con-pho-chuyen-ban-du-mon-tu-ga-ve-dem-o-ha-noi-20180601032149262rf2018073115533278.chn%22%2C%22kichban%22%3A%5B%5D%2C%22tags%22%3A%5B%5D%2C%22payload%22%3A%22keyGgtXMaBR1533112200156%22%2C%22app_facebook_url%22%3A%22%22%2C%22app_fb_attributes%22%3A%5B%22%22%5D%7D%5D%2C%22isCurrent%22%3Atrue%2C%22url%22%3A%22http%3A%2F%2Fkenh14.vn%2Fme-an-thit-ga-nhat-dinh-dung-bo-qua-con-pho-chuyen-ban-du-mon-tu-ga-ve-dem-o-ha-noi-20180601032149262rf2018073115533278.chn%22%2C%22image%22%3A%22https%3A%2F%2Fkenh14cdn.com%2Fzoom%2F600_315%2F2018%2F6%2F1%2Fphoto1527794591355-15277945913561800514566.jpg%22%7D%2C%7B%22id%22%3A4%2C%22title%22%3A%22M%C3%B3n%20%C4%83n%20nh%C3%ACn%20kh%C3%A1%20quen%20thu%E1%BB%99c%20nh%C6%B0ng%20l%E1%BA%A1i%20khi%E1%BA%BFn%20Hye%20Rin%2C%20Jeong%20Hwa%20%28EXID%29%20khen%20ngo...%22%2C%22content%22%3A%22%22%2C%22descriptions%22%3A%22Nh%C3%ACn%20t%C6%B0%E1%BB%9Fng%20m%C3%B3n%20th%E1%BB%8Bt%20n%C6%B0%E1%BB%9Bng%20H%C3%A0n%20Qu%E1%BB%91c%2C%20c%C3%B3%20ng%C6%B0%E1%BB%9Di%20n%C3%B3i%20gi%E1%BB%91ng%20l%E1%BA%A9u%20Shabu%20Shabu%20-%20Nh%E1%BA%ADt...%22%2C%22buttonList%22%3A%5B%5D%2C%22isCurrent%22%3Atrue%2C%22url%22%3A%22http%3A%2F%2Fkenh14.vn%2Fmon-an-nhin-kha-quen-thuoc-nhung-lai-khien-hye-rin-jeong-hwa-exid-khen-ngon-khong-ngot-loi-that-ra-co-gi-dac-biet-20180731234838561.chn%22%2C%22image%22%3A%22https%3A%2F%2Fkenh14cdn.com%2Fzoom%2F600_315%2F2018%2F7%2F31%2Fphoto1533054333523-153305433352359348626.jpg%22%7D%5D"
	tg := s1[:2]
	if tg == "[{" {
		urldecodeStep, _ := url.QueryUnescape(s1)
		fmt.Println("urldecodeStep ", urldecodeStep)
		fmt.Printf("=====1 type %T", urldecodeStep)
		contenJson, _ := Json_decodeArray(urldecodeStep)
		for _, idx := range contenJson {
			if idx["title"] == "" {
				idx["title"] = "abc la abc"
				fmt.Println("\ncontenJson ", idx)
				fmt.Printf(" ===== type %T", idx)
				stepJson, _ := Json_encode(idx)
				fmt.Println("\nstepJson ", stepJson)
				fmt.Printf("type %T", stepJson)
				stepJsonEncode := UrlEncode(stepJson)
				fmt.Println("\nstepJsonEncode ", stepJsonEncode)
				fmt.Printf("type %T", stepJsonEncode)
			}

		}

	}
	if tg == "{\"" {
		urldecodeStep, _ := url.QueryUnescape(s1)
		fmt.Println("***urldecodeStep ", urldecodeStep)
		fmt.Printf("=====1 type %T", urldecodeStep)
		contenJson, _ := Json_decode(urldecodeStep)
		contenJson["title"] = "abc la abc"
		fmt.Println("\ncontenJson ", contenJson)
		fmt.Printf(" ===== type %T", contenJson)
		stepJson, _ := Json_encode(contenJson)
		fmt.Println("\nstepJson ", stepJson)
		fmt.Printf("type %T", stepJson)
		stepJsonEncode := UrlEncode(stepJson)
		fmt.Println("\nstepJsonEncode ", stepJsonEncode)
		fmt.Printf("type %T", stepJsonEncode)
	}

}
func Json_decode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}
func Json_decodeArray(data string) ([]map[string]interface{}, error) {
	var dat []map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}
func Json_encode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

type T struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Image        string `json:"image"`
	Url          string `json:"url"`
	Descriptions string `json:"descriptions"`
	ButtonList   []struct {
		Id                 int           `json:"id"`
		Title              string        `json:"title"`
		Type               string        `json:"type"`
		Content            string        `json:"content"`
		Kichban            []interface{} `json:"kichban"`
		Tags               []interface{} `json:"tags"`
		Payload            string        `json:"payload"`
		AppFacebookUrl     string        `json:"app_facebook_url"`
		AppFbAttributes    []string      `json:"app_fb_attributes"`
		IsShowScriptSelect bool          `json:"isShowScriptSelect,omitempty"`
	} `json:"buttonList"`
	Content   string `json:"content,omitempty"`
	IsCurrent bool   `json:"isCurrent,omitempty"`
}
type T2 struct {
	Title      string `json:"title"`
	ButtonList []struct {
		Id      int    `json:"id"`
		Title   string `json:"title"`
		Type    string `json:"type"`
		Content string `json:"content"`
		Kichban []struct {
			ScriptId   string `json:"script_id"`
			ScriptName string `json:"script_name"`
		} `json:"kichban"`
		Tags               []interface{} `json:"tags"`
		Payload            string        `json:"payload"`
		AppFacebookUrl     string        `json:"app_facebook_url"`
		AppFbAttributes    []string      `json:"app_fb_attributes"`
		IsShowScriptSelect bool          `json:"isShowScriptSelect"`
	} `json:"buttonList"`
}
