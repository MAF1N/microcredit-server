package dal

const ConnectionString = "MAF1N:33726@/course_work"
const Driver = "mysql"

type UnitOfWork struct {
	ClientRepository ClientRepository
	OrganizationRepository OrganizationRepository
	BankRepository BankRepository
	CreditAccountRepository  CreditAccountRepository
	AdminRepository AdminRepository
}