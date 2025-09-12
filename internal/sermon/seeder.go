package sermon

import (
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Cek apakah sudah ada data
	var count int64
	db.Model(&Sermon{}).Count(&count)
	if count > 0 {
		return nil // skip kalau sudah ada data
	}

	sermons := []Sermon{
		{
			Title:     "Jangan Mau Hidup Dihimpit Dan Dipenjara Oleh Tipu Muslihat Iblis",
			MainVerse: "Yesaya 42:21-22",
			Preacher:  "Pdm. Dens Mamesah",
			Date:      time.Date(2025, time.August, 24, 0, 0, 0, 0, time.Local),
			Content: `
		Firman Tuhan ketika datang bukanlah secara kebetulan melainkan karena kebutuhan gereja Tuhan, kebutuhan bangsa Israel saat itu, dan kebutuhan umat Tuhan masa sekarang. Namun, bangsa Israel ketika kebenaran disampaikan mereka pura-pura tuli dan buta, mereka tidak mau mendengarkan Firman Allah sehingga masuk dalam geronggang-geronggang, terpenjara. Damai sejahtera dan sukacita mereka dicuri sehingga mereka tidak dapat berbuat apa-apa. Apa yang terjadi pada zaman Yesaya bisa saja terjadi pada zaman sekarang. Jika kita yang hidup di masa sekarang pura-pura buta dan tuli maka kita tahu apa yang akan terjadi seperti pada bangsa Israel. Iblis berusaha keras untuk mencuri, membunuh, membinasakan, dan karena itu kita di akhir zaman ini harus berusaha keras untuk melawan tipu muslihat iblis supaya kita dapat berdiri sebagai gereja yang imani di hadapan Tuhan.
		
		Firman Tuhan berkata bahwa persembahkanlah tubuhmu sebagai persembahan yang hidup, yang kudus dan berkenan kepada Allah. Itu adalah ibadah yang sejati. Esensi dari ibadah adalah Roma 12:1, dan pada ayat 2, berubahlah oleh pembaharuan budimu, maka kita mampu membedakan mana yang benar dan yang berkenan pada Allah. Yohanes 14:15, "Jika kamu mengasihi Aku, kamu akan menuruti segala perintah-Ku." Ketika kita mengatakan mengasihi Tuhan maka kita harus melakukan perintah-perintah-Nya, supaya hidup kita berkenan kepada-Nya. Yohanes 10:10, "Pencuri datang hanya untuk mencuri dan membunuh dan membinasakan; Aku datang, supaya mereka mempunyai hidup, dan mempunyainya dalam segala kelimpahan." Jika iblis mencuri sesuatu yang bernilai kekal, yaitu firman Allah dalam hati kita, maka tidak heran meskipun telah berdoa dan beribadah tapi kita masih menyimpan kepahitan.
		
		Saat ini kita masih diberikan kesempatan oleh Tuhan untuk mendengarkan firman-Nya, supaya kita tidak hidup terpenjara. Lukas 8:14, "Yang jatuh dalam semak duri ialah orang yang telah mendengar firman itu, dan dalam pertumbuhan selanjutnya mereka terhimpit oleh kekuatiran dan kekayaan dan kenikmatan hidup, sehingga mereka tidak menghasilkan buah yang matang." Mereka terhimpit oleh semak duri sehingga kehidupan kerohanian tidak bertumbuh. Kenapa tidak bertumbuh? Karena ada kekuatiran. Mengapa kuatir? Karena tidak percaya dengan kebenaran firman Tuhan. Kelihatan aktif dalam gereja tetapi tidak mau melakukan perintah Tuhan. Padahal Tuhan mencari orang yang melakukan firman-Nya. Ada orang yang kelihatan aktif melayani tetapi dosa jalan terus. Tetapi sebagai umat Tuhan, kita percaya bahwa kita adalah orang-orang yang mendengar dan melakukan firman Tuhan.
		
		Ketika kita beriman dan melakukan firman Tuhan, mujizat menjadi nyata, janji Tuhan menjadi nyata dalam hidup kita. Tuhan tidak pernah mengingkari janji-Nya. Bangsa Israel mendengar Firman Tuhan tetapi tidak melakukan. Kita janganlah hidup demikian. Kita mendengar tapi tidak melakukan, berpikir lebih tinggi dari Firman Tuhan, dan mengandalkan diri sendiri. Akhirnya iblis mencuri damai sejahtera dan sukacita. Kehidupan kita haruslah dipenuhi damai sejahtera dan sukacita yang dari Allah. Dan untuk mendapatkannya yaitu dengan mendengar dan melakukan Firman Tuhan.
		
		Dalam Kejadian 13 diceritakan mengenai Abram, yang adalah orang yang percaya dan beriman pada Tuhan. Abram pun adalah orang yang memiliki harta dan ternak yang banyak. Lot keponakannya pun demikian. Namun Lot memiliki sifat iri kepada Abram sehingga dengan hikmat Tuhan terjadilah percakapan antara Abram dan Lot agar tidak terjadi perkelahian di antara mereka. Dalam hal ini Abram berperang dengan iblis. Dan kemudian Allah memberkati Abram. Sasaran iblis adalah pikiran manusia. Karena itu isilah pikiran kita dengan Firman Tuhan dan kita harus mampu mendeteksi tipu muslihat iblis. Tipu muslihat iblis adalah merampas sesuatu dalam pikiran kita. Kita harus menjadi orang-orang yang memiliki pikiran Kristus.
		
		Seperti Abraham, ia penuh belas kasihan sehingga mampu berkata pada Lot untuk memilih tanah mana yang Lot mau di daerah yang begitu luas itu. Abraham melakukan hal tersebut agar tidak terjadi perkelahian di antara mereka. Kitapun harus seperti Abraham. Jika kita tidak bisa memiliki pikiran Kristus, kita akan menjadi sumber keributan. Orang-orang yang demikianlah yang tuli dan buta.
		
		Lot adalah saudara dari Abraham. Abraham ketika melihat akan terjadi pertikaian, ia bertindak untuk memisahkan diri. Kita adalah gereja Tuhan, sehingga kita harus memisahkan diri dari dunia. Galatia 5:16, "Maksudku ialah: hiduplah oleh Roh, maka kamu tidak akan menuruti keinginan daging." Supaya hidup tidak seperti dalam Yesaya 42:21-22, yang terpenjara, terkurung, terhimpit oleh semak, maka kita harus dipenuhi oleh Roh Kudus. Dengan Roh Kudus kita akan mampu melakukan perintah-perintah Tuhan. Dan dengan demikian kita akan berjalan terus sampai pada kekekalan. Bukalah hati, tangkaplah Firman Tuhan, Roh Kudus hadir dan mengeluarkan nafsu-nafsu dunia sehingga kita hidup bebas merdeka bersama dengan Tuhan. Amin.
		`,
			ImageURL:     "https://example.com/images/sermon1.jpg",
			ProfileImage: "https://example.com/images/preacher1.jpg",
		},
		{
			Title:     "Membalas Kebaikan Tuhan",
			MainVerse: "Mikha 6:8",
			Preacher:  "Pdm. Marky A. Lumanauw, SH",
			Date:      time.Date(2025, time.August, 31, 0, 0, 0, 0, time.Local),
			Content: `
			Melalui FirmanNya dalam Mikha 6:8, “Hai manusia, telah diberitahukan kepadamu, apa yang baik. Dan apakah yang dituntut TUHAN dari padamu: selain berlaku adil, mencintai kesetiaan, dan hidup dengan rendah hati di hadapan Allahmu” Tuhan telah memberitahukan apa yang baik untuk kita umatNya.
			
			Kenapa Tuhan memberitahukan hal ini? Karena Tuhan tahu kecenderungan hati manusia untuk berbuat jahat. Seperti pada Kejadian 6:5, "Ketika dilihat TUHAN, bahwa kejahatan manusia besar di bumi dan bahwa segala kecenderungan hatinya selalu membuahkan kejahatan semata-mata." Dan karena manusia cenderung berbuat jahat maka Allah pernah menyesal telah menciptakan manusia. Kejadian 6:6, "Maka menyesallah TUHAN, bahwa Ia telah menjadikan manusia di bumi, dan hal itu memilukan hati-Nya."
			
			... (paragraf pengantar) ...
			
			Menurut Mikha 6:8, imbalan yang harus kita berikan adalah:
			
			1. **Berlaku adil di hadapan-Nya**  
			   Berlaku adil artinya tidak berat sebelah, seimbang. Ibrani 1:8-9, "Tetapi tentang Anak Ia berkata: Takhta-Mu, ya Allah, tetap untuk seterusnya dan selamanya, dan tongkat kerajaan-Mu adalah tongkat kebenaran..."  
			
			2. **Mencintai kesetiaan di hadapan Allah**  
			   Setia artinya tetap berpegang teguh, tidak pernah ingkar. Contoh: setia memberi perpuluhan, setia dalam pelayanan, setia berkorban, setia melakukan kehendak Tuhan. Dalam terjemahan KJV dipakai kata *to love mercy* yang berarti juga mengasihi. Kalau kita hidup dalam kasih, kita akan bertemu dengan 1 Korintus 13:4-7.  
			
			3. **Hidup dengan rendah hati**  
			   Kita hidup di dunia ini dilihat oleh Tuhan di sorga. Henokh adalah contoh orang yang rendah hati, yang hidup bergaul karib dengan Allah. Orang yang bergaul dengan Allah adalah orang yang berjalan dengan Allah. Tuhan tidak melihat aktivitas gereja kita, tapi apakah kita berlaku adil, mencintai kesetiaan, dan hidup rendah hati di hadapan-Nya.  
			   
			Amin.
			`,
			ImageURL:     "https://example.com/images/sermon2.jpg",
			ProfileImage: "https://example.com/images/preacher2.jpg",
		},
		{
			Title:        "Pengharapan yang Teguh",
			MainVerse:    "Yeremia 29:11",
			Preacher:     "Pdt. Daniel",
			Date:         time.Now().AddDate(0, 0, -10),
			Content:      "Isi khotbah mengenai janji Tuhan yang memberi masa depan penuh harapan.",
			ImageURL:     "https://example.com/images/sermon3.jpg",
			ProfileImage: "https://example.com/images/preacher3.jpg",
		},
		{
			Title:        "Kuasa dalam Doa",
			MainVerse:    "Yakobus 5:16",
			Preacher:     "Pdt. Lydia",
			Date:         time.Now().AddDate(0, 0, -15),
			Content:      "Isi khotbah mengenai doa yang penuh kuasa dan efektif.",
			ImageURL:     "https://example.com/images/sermon4.jpg",
			ProfileImage: "https://example.com/images/preacher4.jpg",
		},
		{
			Title:        "Kasih yang Mengubahkan",
			MainVerse:    "Yohanes 3:16",
			Preacher:     "Pdt. Andreas Rawung",
			Date:         time.Date(2025, 8, 20, 10, 0, 0, 0, time.UTC),
			Content:      "Isi khotbah tentang kasih Allah yang mengubahkan hidup...",
			ImageURL:     "https://picsum.photos/seed/2/600/400",
			ProfileImage: "https://randomuser.me/api/portraits/women/44.jpg",
		},
		{
			Title:        "Hidup Berkenan kepada Tuhan",
			MainVerse:    "Roma 12:1-2",
			Preacher:     "Pdt. Lydia Lumintang",
			Date:         time.Date(2025, 7, 15, 10, 0, 0, 0, time.UTC),
			Content:      "Isi khotbah tentang persembahan hidup yang berkenan...",
			ImageURL:     "https://picsum.photos/seed/3/600/400",
			ProfileImage: "https://randomuser.me/api/portraits/men/46.jpg",
		},
		{
			Title:        "Kuasa Doa dalam Penderitaan",
			MainVerse:    "Yakobus 5:13-16",
			Preacher:     "Pdt. Sarah Saranesa",
			Date:         time.Date(2025, 9, 5, 10, 0, 0, 0, time.UTC),
			Content:      "Isi khotbah tentang kekuatan doa ketika dalam penderitaan...",
			ImageURL:     "https://picsum.photos/seed/4/600/400",
			ProfileImage: "https://randomuser.me/api/portraits/women/47.jpg",
		},
	}

	return db.Create(&sermons).Error
}
