package lazada

type GetSellerResponse struct {
	NameCompany string `json:"name_company"`
	Name        string `json:"name"`
	Verified    bool   `json:"verified"`
	SellerId    int64  `json:"seller_id"`
	Email       string `json:"email"`
	ShortCode   string `json:"short_code"`
	Cb          bool   `json:"cb"`
	Status      string `json:"status"`
}

type GetOrderDetailResponse []OrderDetailResponse

type GetOrderResponse OrderResponse
