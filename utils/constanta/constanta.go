package constanta

// Constanta For Role
const (
	USER       = "user"
	ADMIN      = "admin"
	SUPERADMIN = "super_admin"
)

const (
	URL_STORAGE = "https://cimxqffotlogzqvadisz.supabase.co/storage/v1"
	URL         = URL_STORAGE + "/object/public/"
)


// Constanta For Success
const (
	SUCCESS_LOGIN       = "berhasil melakukan login"
	SUCCESS_NULL        = "data belum tersedia"
	SUCCESS_CREATE_DATA = "berhasil membuat data"
	SUCCESS_DELETE_DATA = "berhasil menghapus data"
	SUCCESS_GET_DATA    = "berhasil mendapatkan data"
)

// Constanta For Utils
const (
	//VERIFICATION_URL = "http://localhost:8080/verify-token?token="
	VERIFICATION_URL   = "https://api.recything.my.id/verify-token?token="
	EMAIL_NOT_REGISTER = "email belum terdaftar"
	IMAGE_ADMIN        = "https://ui-avatars.com/api/?background=56cc33&color=fff&name="
)

// Constanta For Error
const (
	ERROR_TEMPLATE         = "gagal menguraikan template"
	ERROR_DATA_ID          = "id tidak ditemukan"
	ERROR_ID_INVALID       = "id salah"
	ERROR_DATA_EMAIL       = "email tidak ditemukan"
	ERROR_FORMAT_EMAIL     = "error : format email tidak valid"
	ERROR_EMAIL_EXIST      = "error : email sudah digunakan"
	ERROR_AKSES_ROLE       = "akses ditolak"
	ERROR_PASSWORD         = "error : password lama tidak sesuai"
	ERROR_CONFIRM_PASSWORD = "error : konfirmasi password tidak sesuai"
	ERROR_EXTRA_TOKEN      = "gagal ekstrak token"
	ERROR_ID_ROLE          = "id atau role tidak ditemukan"
	ERROR_GET_DATA         = "data tidak ditemukan"
	ERROR_EMPTY            = "error : harap lengkapi data dengan benar"
	ERROR_EMPTY_FILE       = "error : tidak ada file yang di upload"
	ERROR_HASH_PASSWORD    = "error : hash password"
	ERROR_DATA_NOT_FOUND   = "data tidak ditemukan"
	ERROR_DATA_EXIST       = "error : data sudah ada"
	ERROR_MISSION_LIMIT    = "error : tahapan misi tidak boleh dari 5"
	ERROR_INVALID_TITLE    = "error: tahapan misi tidak boleh sama"
	ERROR_INVALID_ID       = "error: id tidak boleh sama"
	ERROR_INVALID_UPDATE   = "error: data harus berberbeda dengan data sebelumnya"
	ERROR_INVALID_INPUT    = "data yang diinput tidak sesuai"
	ERROR_NOT_FOUND        = "data tidak ditemukan"
	ERROR_RECORD_NOT_FOUND = "record not found"
	ERROR_INVALID_TYPE     = "berupa angka"
	ERROR_INVALID_STATUS   = "status tidak valid"
	ERROR_LIMIT            = "error : limit tidak boleh lebih dari 10"
	ERROR_LENGTH_PASSWORD  = "error : minimal 8 karakter, ulangi kembali"
)

// Message Handle
const (
	ALREADY = "sudah"
	NO      = "tidak"
	MUST    = "harus"
	FAILED  = "gagal"
	ERROR   = "error"
)

// const for fix data
var (
	Unit                = []string{"barang", "kilogram"}
	Category            = []string{"sampah anorganik", "sampah organik", "informasi", "batasan"}
	Days                = []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	REPORT_TYPE         = []string{"pelanggaran sampah", "tumpukan sampah"}
	SCALE_TYPE          = []string{"skala besar", "skala kecil"}
	STATUS_EVENT        = []string{"berjalan", "belum berjalan", "selesai"}
	Status_Exchange     = []string{"diproses", "selesai"}
	STATUS_MISSION_USER = []string{"berjalan", "selesai"}
	STATUS_ADMIN        = []string{"aktif", "tidak aktif"}
	CATEGORY_ARTICLE    = []string{"plastik", "kaca", "logam", "organik", "kertas", "kaleng", "minyak", "elektronik", "tekstil", "baterai"}
	ERROR_MESSAGE       = []string{"sudah", "tidak", "harus", "gagal", "harap"}
	FILTER_PROMPT       = []string{ORGANIC, ANORGANIC, LIMITATION, INFORMATION}
)

// const for mission
const (
	OVERDUE   = "Melewati Tenggat"
	ACTIVE    = "Aktif"
	MAX_STAGE = 1
	MIN_STAGE = 1
)

// const status
const (
	PERLU_TINJAUAN = "perlu tinjauan"
	DISETUJUI      = "disetujui"
	DITOLAK        = "ditolak"
	DIPROSES       = "diproses"
	SELESAI        = "selesai"
	TERBARU        = "terbaru"
	BERJALAN       = "berjalan"
	GAGAL          = "gagal"
)

const (
	NeedProof  = "upload bukti pengerjaan"
	NeedReview = "menunggu verifikasi"
)

const (
	BRONZE   = "bronze"
	PLATINUM = "platinum"
	SILVER   = "silver"
	GOLD     = "gold"
)

const (
	ASIABANGKOK = "Asia/Bangkok"
)

const (
	ORGANIC     = "sampah organik"
	ANORGANIC   = "sampah anorganik"
	INFORMATION = "informasi"
	LIMITATION  = "batasan"
)
