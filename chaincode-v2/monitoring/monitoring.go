package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//External Chaincode Server Config on main function
type serverConfig struct {
	CCID    string
	Address string
}

//Smart Contract for Monitoring Aloptama(Baku) + AlatOto(ARG&AWS)
type MonitoringContract struct {
	contractapi.Contract
}

// Aloptama declare data Alat Operasional
type Aloptama struct {
	KodeAlat     string `json:"kodealat"` // for legal notes
	NamaAlat     string `json:"namaalat"`
	MerekAlat    string `json:"merekalat"`
	JumlahAlat   int64  `json:"jumlahalat"`
	ThnPengadaan int64  `json:"thnpengadaan"`
	Kondisi      string `json:"kondisi"`
	Keterangan   string `json:"keterangan"`
	TxID         string `json:"txId"`
	CreatedAt    int64  `json:"createdAt"`
}

// AlatOto declare data Alat Otomatis
type AlatOto struct {
	KodeSite    string  `json:"kodesite"` // for legal notes
	NamaSite    string  `json:"namasite"`
	JenisAlat   string  `json:"jenisalat"`
	LokasiSite  string  `json:"lokasisite"`
	MerekSensor string  `json:"mereksensor"`
	MerekLogger string  `json:"mereklogger"`
	ResSensor   float64 `json:"ressensor"`
	KapBaterai  int64   `json:"kapbaterai"`
	KapSolar    int64   `json:"kapsolar"`
	CorrMT      string  `json:"corrmt"`
	PrevMT      string  `json:"prevmt"`
	TxID        string  `json:"TxID"`
	CreatedAt   int64   `json:"createdAt"`
}

//HistoryAloptamaResult structure used for returning result of history query
type HistoryAloptamaResult struct {
	Record    *Aloptama `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
}

//HistoryAlatOtoResult structure used for returning result of history query
type HistoryAlatOtoResult struct {
	Record    *AlatOto  `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
}

//QueryAloptamaResult structure used for handling result of query
type QueryAloptamaResult struct {
	Key    string `json:"key"`
	Record *Aloptama
}

//QueryAlatOtoResult structure used for handling result of query
type QueryAlatOtoResult struct {
	Key    string `json:"key"`
	Record *AlatOto
}

// Create Aloptama Asset to world state
func (s *MonitoringContract) CreateAloptama(ctx contractapi.TransactionContextInterface, kodealat, namaalat, merekalat string, jumlahalat, tahunpengadaan int64, kondisi, keterangan string) error {
	exists, err := s.AloptamaExists(ctx, kodealat)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("aloptama asset %s already exist", kodealat)
	}

	aloptama := Aloptama{
		KodeAlat:     kodealat,
		NamaAlat:     namaalat,
		MerekAlat:    merekalat,
		JumlahAlat:   jumlahalat,
		ThnPengadaan: tahunpengadaan,
		Kondisi:      kondisi,
		Keterangan:   keterangan,
		TxID:         ctx.GetStub().GetTxID(),
		CreatedAt:    time.Now().Unix(),
	}

	aloptamaJSON, err := json.Marshal(aloptama)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(kodealat, aloptamaJSON)
}

//CreateAlatOto create AlatOto asset to ledger
func (s *MonitoringContract) CreateAlatOto(ctx contractapi.TransactionContextInterface, kodesite, namasite, jenisalat, lokasisite, mereksensor, mereklogger string, ressensor float64, kapbaterai, kapsolar int64, corrmt, prevmt string) error {
	exists, err := s.AlatOtoExists(ctx, kodesite)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("alat otomatis asset %s already exist", kodesite)
	}

	alatoto := AlatOto{
		KodeSite:    kodesite,
		NamaSite:    namasite,
		JenisAlat:   jenisalat,
		LokasiSite:  lokasisite,
		MerekSensor: mereksensor,
		MerekLogger: mereklogger,
		ResSensor:   ressensor,
		KapBaterai:  kapbaterai,
		KapSolar:    kapsolar,
		CorrMT:      corrmt,
		PrevMT:      prevmt,
		TxID:        ctx.GetStub().GetTxID(),
		CreatedAt:   time.Now().Unix(),
	}

	alatotoJSON, err := json.Marshal(alatoto)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(kodesite, alatotoJSON)
}

//ReadAloptama return the aloptama asset stored in world state with given kodealat (id)
func (s *MonitoringContract) ReadAloptama(ctx contractapi.TransactionContextInterface, kodealat string) (*Aloptama, error) {
	aloptamaJSON, err := ctx.GetStub().GetState(kodealat)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}
	if aloptamaJSON == nil {
		return nil, fmt.Errorf("the aloptama %s asset does not exist", kodealat)
	}

	var aloptama Aloptama
	err = json.Unmarshal(aloptamaJSON, &aloptama)

	if err != nil {
		return nil, err
	}

	return &aloptama, nil
}

//ReadAlatOto return the alatoto asset stored in world state with given kodealat (id)
func (s *MonitoringContract) ReadAlatOto(ctx contractapi.TransactionContextInterface, kodesite string) (*AlatOto, error) {
	alatotoJSON, err := ctx.GetStub().GetState(kodesite)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}
	if alatotoJSON == nil {
		return nil, fmt.Errorf("the alatoto %s asset does not exist", kodesite)
	}

	var alatoto AlatOto
	err = json.Unmarshal(alatotoJSON, &alatoto)

	if err != nil {
		return nil, err
	}

	return &alatoto, nil
}

/*UpdateAloptama update all existing aloptama asset datas in the world state with provided parameters
func (s *MonitoringContract) UpdateAloptama(ctx contractapi.TransactionContextInterface, kodealat, namaalat, merekalat string, jumlahalat, tahunpengadaan int64, kondisi, keterangan string) (*Aloptama, error) {
	aloptamaJSON, err := ctx.GetStub().GetState(kodealat)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}
	if aloptamaJSON == nil {
		return nil, fmt.Errorf("the alatoto %s asset does not exist", kodealat)
	}

	var aloptama Aloptama

	aloptama.TxID = ctx.GetStub().GetTxID()
	aloptama.CreatedAt = time.Now().Unix()

	err = json.Unmarshal(aloptamaJSON, &aloptama)
	if err != nil {
		return nil, err
	}

	return &aloptama, nil
}

//UpdateAloptama update all existing aloptama asset datas in the world state with provided parameters
func (s *MonitoringContract) UpdateAlatOto(ctx contractapi.TransactionContextInterface, kodesite, namasite, jenisalat, lokasisite, mereksensor, mereklogger string, ressensor float64, kapbaterai, kapsolar int64, corrmt, prevmt string) (*AlatOto, error) {
	alatotoJSON, err := ctx.GetStub().GetState(kodesite)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}
	if alatotoJSON == nil {
		return nil, fmt.Errorf("the alatoto %s asset does not exist", kodesite)
	}

	var alatoto AlatOto

	alatoto.TxID = ctx.GetStub().GetTxID()
	alatoto.CreatedAt = time.Now().Unix()

	err = json.Unmarshal(alatotoJSON, &alatoto)
	if err != nil {
		return nil, err
	}

	return &alatoto, nil
}*/

//--Check Existence in world state--//

//AloptamaExists check Aloptama existence in world state
func (s *MonitoringContract) AloptamaExists(ctx contractapi.TransactionContextInterface, kodealat string) (bool, error) {
	aloptamaJSON, err := ctx.GetStub().GetState(kodealat)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	return aloptamaJSON != nil, nil
}

//AlatOtoExists check AlatOto existence in world state
func (s *MonitoringContract) AlatOtoExists(ctx contractapi.TransactionContextInterface, kodesite string) (bool, error) {
	alatotoJSON, err := ctx.GetStub().GetState(kodesite)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	return alatotoJSON != nil, nil
}

// --Update Kondisi+Keterangan Field in Aloptama data-- //
func (s *MonitoringContract) UpdateKondisiAloptama(ctx contractapi.TransactionContextInterface, kodealat string, newKondisi string, newKeterangan string) error {
	aloptama, err := s.ReadAloptama(ctx, kodealat)
	if err != nil {
		return err
	}

	aloptama.Kondisi = newKondisi
	aloptama.Keterangan = newKeterangan
	aloptama.TxID = ctx.GetStub().GetTxID()
	aloptama.CreatedAt = time.Now().Unix()

	aloptamaJSON, err := json.Marshal(aloptama)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(kodealat, aloptamaJSON)
}

//--Update PM+CM Field in AlatOto--//
func (s *MonitoringContract) UpdatePMCM(ctx contractapi.TransactionContextInterface, kodesite string, newPrevMT string, newCorrMT string) error {
	alatoto, err := s.ReadAlatOto(ctx, kodesite)
	if err != nil {
		return err
	}

	alatoto.PrevMT = newPrevMT
	alatoto.CorrMT = newCorrMT
	alatoto.TxID = ctx.GetStub().GetTxID()
	alatoto.CreatedAt = time.Now().Unix()

	alatotoJSON, err := json.Marshal(alatoto)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(kodesite, alatotoJSON)
}

//--Delete Feature (beta)--//

//DeleteAloptama deletes a given aloptama asset from the world state
func (s *MonitoringContract) DeleteAloptama(ctx contractapi.TransactionContextInterface, kodealat string) error {
	exists, err := s.AloptamaExists(ctx, kodealat)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the aloptama %s asset does not exist", kodealat)
	}

	return ctx.GetStub().DelState(kodealat)
}

//DeleteAloptama deletes a given aloptama asset from the world state
func (s *MonitoringContract) DeleteAlatOto(ctx contractapi.TransactionContextInterface, kodesite string) error {
	exists, err := s.AlatOtoExists(ctx, kodesite)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the alatoto %s asset does not exist", kodesite)
	}

	return ctx.GetStub().DelState(kodesite)
}

func main() {
	// See chaincode.env.example
	config := serverConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	chaincode, err := contractapi.NewChaincode(new(MonitoringContract))

	if err != nil {
		fmt.Printf("Error create Aloptama+AlatOto Chaincode: %s", err.Error())
		return
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting chaincodes: %s", err.Error())
	}

}
