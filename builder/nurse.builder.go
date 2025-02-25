package builder

import (
	"github.com/jmoiron/sqlx"

	nurseexternalrpc "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/infars/externalrpc"
	nurserepository "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/infars/repository"
	nursecommands "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/commands"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
)

// hehe
type builderOfNurse struct {
	db                    *sqlx.DB
	urlPathAccountService string
}

func NewNurseBuilder(db *sqlx.DB) builderOfNurse {
	return builderOfNurse{db: db}
}

func (s builderOfNurse) AddUrlPathAccountService(url string) builderOfNurse {
	s.urlPathAccountService = url
	return s
}

func (s builderOfNurse) BuildNurseCmdRepo() nursecommands.NurseCommandRepo {
	return nurserepository.NewNurseRepo(s.db)
}

func (s builderOfNurse) BuildExternalAccountServiceInCmd() nursecommands.ExternalAccountService {
	return nurseexternalrpc.NewAccountRPC(s.urlPathAccountService)
}

func (s builderOfNurse) BuildNurseQueryRepo() nursequeries.NurseQueryRepo {
	return nurserepository.NewNurseRepo(s.db)
}

func (s builderOfNurse) BuildExternalAccountServiceInQuery() nursequeries.ExternalAccountService {
	return nurseexternalrpc.NewAccountRPC(s.urlPathAccountService)
}
