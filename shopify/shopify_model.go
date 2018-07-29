package shopify

type ShopifyAuthResponse struct {
	AccessToken         string `json:"access_token"`
	Scope               string `json:"scope"`
	ExpiresIn           int    `json:"expires_in"`
	AssociatedUserScope string `json:"associated_user_scope"`
	AssociatedUser      struct {
		ID           int    `json:"id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		AccountOwner bool   `json:"account_owner"`
	} `json:"associated_user"`
}

type ShopifyApiProductCount struct {
	Count int `json:"count"`
}

type ShopifyScriptTagsResponse struct {
	ScriptTags []struct {
		ID           int    `json:"id"`
		Src          string `json:"src"`
		Event        string `json:"event"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		DisplayScope string `json:"display_scope"`
	} `json:"script_tags"`
}

type ShopifyPostScriptTagResponse struct {
	ScriptTag struct {
		ID           int    `json:"id"`
		Src          string `json:"src"`
		Event        string `json:"event"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		DisplayScope string `json:"display_scope"`
	} `json:"script_tag"`
}

type ShopifyClientConfig struct {
	Shop          string                 `json:"shop"`
	IndexAddress  string                 `json:"indexAddress"`
	AuthResponse  ShopifyAuthResponse    `json:"shopifyAuthResponse"`
	SearchConfigs []*ShopifySearchConfig `json:"shopifySearchConfigs"`
}

type ShopifySearchConfig struct {
	Name                      string `json:"name" schema:"Name"`
	IsActive                  bool   `json:"isActive" schema:"IsActive"`
	IncludeProductSuggesitons bool   `json:"includeProductSuggesitons" schema:"IncludeProductSuggestions"`
}

type ShopifyApiProductsResponse struct {
	Products []struct {
		ID                int         `json:"id"`
		Title             string      `json:"title"`
		BodyHTML          string      `json:"body_html"`
		Vendor            string      `json:"vendor"`
		ProductType       string      `json:"product_type"`
		CreatedAt         string      `json:"created_at"`
		Handle            string      `json:"handle"`
		UpdatedAt         string      `json:"updated_at"`
		PublishedAt       string      `json:"published_at"`
		TemplateSuffix    interface{} `json:"template_suffix"`
		Tags              string      `json:"tags"`
		PublishedScope    string      `json:"published_scope"`
		AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
		Variants          []struct {
			ID                   int         `json:"id"`
			ProductID            int         `json:"product_id"`
			Title                string      `json:"title"`
			Price                string      `json:"price"`
			Sku                  string      `json:"sku"`
			Position             int         `json:"position"`
			InventoryPolicy      string      `json:"inventory_policy"`
			CompareAtPrice       interface{} `json:"compare_at_price"`
			FulfillmentService   string      `json:"fulfillment_service"`
			InventoryManagement  string      `json:"inventory_management"`
			Option1              string      `json:"option1"`
			Option2              interface{} `json:"option2"`
			Option3              interface{} `json:"option3"`
			CreatedAt            string      `json:"created_at"`
			UpdatedAt            string      `json:"updated_at"`
			Taxable              bool        `json:"taxable"`
			Barcode              interface{} `json:"barcode"`
			Grams                int         `json:"grams"`
			ImageID              interface{} `json:"image_id"`
			InventoryQuantity    int         `json:"inventory_quantity"`
			Weight               interface{} `json:"weight"`
			WeightUnit           string      `json:"weight_unit"`
			InventoryItemID      int64       `json:"inventory_item_id"`
			OldInventoryQuantity int         `json:"old_inventory_quantity"`
			RequiresShipping     bool        `json:"requires_shipping"`
			AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
		} `json:"variants"`
		Options []struct {
			ID        int64    `json:"id"`
			ProductID int64    `json:"product_id"`
			Name      string   `json:"name"`
			Position  int      `json:"position"`
			Values    []string `json:"values"`
		} `json:"options"`
		Images []struct {
			ID                int64         `json:"id"`
			ProductID         int64         `json:"product_id"`
			Position          int           `json:"position"`
			CreatedAt         string        `json:"created_at"`
			UpdatedAt         string        `json:"updated_at"`
			Alt               interface{}   `json:"alt"`
			Width             int           `json:"width"`
			Height            int           `json:"height"`
			Src               string        `json:"src"`
			VariantIds        []interface{} `json:"variant_ids"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
		} `json:"images"`
		Image struct {
			ID                int64         `json:"id"`
			ProductID         int64         `json:"product_id"`
			Position          int           `json:"position"`
			CreatedAt         string        `json:"created_at"`
			UpdatedAt         string        `json:"updated_at"`
			Alt               interface{}   `json:"alt"`
			Width             int           `json:"width"`
			Height            int           `json:"height"`
			Src               string        `json:"src"`
			VariantIds        []interface{} `json:"variant_ids"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
		} `json:"image"`
	} `json:"products"`
}

type ShopifyGetThemesResponse struct {
	Themes []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
		Role              string `json:"role"`
		ThemeStoreID      int    `json:"theme_store_id"`
		Previewable       bool   `json:"previewable"`
		Processing        bool   `json:"processing"`
		AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	} `json:"themes"`
}

type ShopifyAssetPutPayload struct {
	Asset struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"asset"`
}
