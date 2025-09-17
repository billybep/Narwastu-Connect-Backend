package organization

import (
	"fmt"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	placeholder := "https://jdrsmpsvpcdbdfjjamja.supabase.co/storage/v1/object/public/NarwastuStorage/Organizations/user_profile_placeholder.jpg"

	organizations := []Organization{
		// 🔹 Pimpinan
		{Category: "Pimpinan", Jabatan: "Gembala Sidang", Nama: "Pdm. Marky A. Lumanauw, S.H", Quote: "Memimpin dengan kasih dan keteladanan.", ProfilePic: placeholder},
		{Category: "Pimpinan", Jabatan: "Wakil Gembala Sidang", Nama: "Pdt. Elsje N. Tangka", Quote: "Melayani sepenuh hati.", ProfilePic: placeholder},
		{Category: "Pimpinan", Jabatan: "Sekretaris Jemaat", Nama: "Yemima Haluang, S.Si", Quote: "Setia dalam pelayanan administrasi.", ProfilePic: placeholder},
		{Category: "Pimpinan", Jabatan: "Bendahara Jemaat", Nama: "Christy Walintukan, S.E", Quote: "Mengelola berkat Tuhan dengan bijaksana.", ProfilePic: placeholder},

		// 🔹 Tua-Tua Sidang
		{Category: "Tua-Tua Sidang", Jabatan: "Tua-tua Sidang", Nama: "Johnny A. Surentu", Quote: "Menjadi teladan bagi jemaat.", ProfilePic: placeholder},
		{Category: "Tua-Tua Sidang", Jabatan: "Tua-tua Sidang", Nama: "Sientje Kaseger-Lontaan", Quote: "Hidup untuk melayani.", ProfilePic: placeholder},

		// 🔹 Staf Penggembalaan
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Bobby S. Rumagit", Quote: "Menggembalakan dengan kerendahan hati.", ProfilePic: placeholder},
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Rosdiana Mamesah", Quote: "Melayani tanpa lelah.", ProfilePic: placeholder},
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Masye Mamesah", Quote: "Kasih Kristus menguatkan pelayanan.", ProfilePic: placeholder},
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Lineke S. Lomboan, S.Th", Quote: "Firman Tuhan menjadi pedoman hidup.", ProfilePic: placeholder},
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Yanny Wewengkang", Quote: "Dipanggil untuk melayani.", ProfilePic: placeholder},
		{Category: "Staf Penggembalaan", Jabatan: "Staf Penggembalaan", Nama: "Pdm. Dens Mamesah", Quote: "Hidup adalah Kristus, mati adalah keuntungan.", ProfilePic: placeholder},

		// 🔹 Komisi-Komisi Jemaat
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelprip", Nama: "Stenly S. Kandouw", Quote: "Bersatu dalam kasih persaudaraan.", ProfilePic: placeholder},
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelwap", Nama: "Pdm. Mery Bolung", Quote: "Melayani kaum wanita dengan setia.", ProfilePic: placeholder},
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelpap", Nama: "Pdm. Alice B. Lumanauw, S.H", Quote: "Menjadi terang bagi generasi muda.", ProfilePic: placeholder},
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelrap", Nama: "Pdm. Stief Roland Kandouw", Quote: "Melayani kaum remaja dengan penuh kasih.", ProfilePic: placeholder},
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelnap", Nama: "Ivonne Kandouw, S.Pd", Quote: "Membentuk generasi anak takut akan Tuhan.", ProfilePic: placeholder},
		{Category: "Komisi-Komisi Jemaat", Jabatan: "Ketua Komisi Pelprup", Nama: "Pdm. Dens Mamesah", Quote: "Melayani kaum pria dengan keteladanan.", ProfilePic: placeholder},

		// 🔹 Rayon-Rayon
		{Category: "Rayon-Rayon", Jabatan: "Ketua Rayon Zion", Nama: "Pdm. Yanny Wewengkang", Quote: "Bertumbuh bersama dalam iman.", ProfilePic: placeholder},
		{Category: "Rayon-Rayon", Jabatan: "Ketua Rayon Karmel", Nama: "Pdm. Rosdiana Mamesah", Quote: "Membangun jemaat dengan kasih.", ProfilePic: placeholder},
		{Category: "Rayon-Rayon", Jabatan: "Ketua Rayon Hermon", Nama: "Pdm. Bobby S. Rumagit", Quote: "Setia dalam doa dan pelayanan.", ProfilePic: placeholder},
		{Category: "Rayon-Rayon", Jabatan: "Ketua Rayon Zaitun", Nama: "Pdm. Masye Mamesah", Quote: "Menghasilkan buah yang baik.", ProfilePic: placeholder},
		{Category: "Rayon-Rayon", Jabatan: "Ketua Rayon Horeb", Nama: "Pdm. Lineke S. Lomboan, S.Th", Quote: "Hidup dalam pengajaran Firman.", ProfilePic: placeholder},

		// 🔹 Wadah-Wadah Pelayanan
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Pelayanan Usia Anugerah", Nama: "Ruddy Lumingkewas", Quote: "Menguatkan iman di masa tua.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Keluarga Senior Imanuel", Nama: "Joan Tumilantouw, S.E", Quote: "Bersama Tuhan sampai akhir.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Keluarga Muda Kasih Agape", Nama: "Betsy Supit", Quote: "Keluarga muda yang hidup dalam kasih.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Diakonia", Nama: "Artje Lumeno", Quote: "Melayani sesama dengan kasih nyata.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Rukun MKM", Nama: "Eiva J. Dien", Quote: "Saling menopang dalam kasih Kristus.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Tika Pro", Nama: "Yulliana M. Mumek", Quote: "Berkarya untuk kemuliaan Tuhan.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua NMC", Nama: "Marcel Saraun, S.E", Quote: "Melayani dengan sukacita.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Rukun Zadok", Nama: "Pdm. Bobby S. Rumagit", Quote: "Taat dan setia dalam pelayanan.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Pelayan Altar", Nama: "Pdm. Mery Bolung", Quote: "Melayani di hadapan Tuhan dengan kudus.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Musik Gereja", Nama: "Billy Pesoth", Quote: "Musik untuk kemuliaan nama Tuhan.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Multi Media", Nama: "Leanditro Andreas Kandouw", Quote: "Melayani Tuhan dengan teknologi.", ProfilePic: placeholder},
		{Category: "Wadah-Wadah Pelayanan", Jabatan: "Ketua Sound System dan Soundman", Nama: "Pdm. Stief Roland Kandouw", Quote: "Mengatur suara bagi kemuliaan Tuhan.", ProfilePic: placeholder},
	}

	for _, org := range organizations {
		var existing Organization
		err := db.Where("category = ? AND jabatan = ? AND nama = ?", org.Category, org.Jabatan, org.Nama).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&org).Error; err != nil {
				return fmt.Errorf("failed to seed organization: %w", err)
			}
		}
	}

	return nil
}
