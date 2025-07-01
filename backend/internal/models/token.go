package models

type BalanceResponse struct {
    Address string `json:"address"`
    Balance string `json:"balance"`
}

type TransferRequest struct {
    From   string `json:"from"`
    To     string `json:"to"`
    Amount string `json:"amount"`
    JWT    string `json:"jwt"`
}

type ApproveRequest struct {
    Owner   string `json:"owner"`
    Spender string `json:"spender"`
    Amount  string `json:"amount"`
    JWT     string `json:"jwt"`
}

type BurnRequest struct {
    Address string `json:"address"`
    Amount  string `json:"amount"`
    JWT     string `json:"jwt"`
}