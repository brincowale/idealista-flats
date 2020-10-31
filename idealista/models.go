package idealista

type Login struct {
	AccessToken string `json:"access_token"`
}

type Results struct {
	Properties     []Property `json:"elementList"`
	Total          int        `json:"total"`
	TotalPages     int        `json:"totalPages"`
	ActualPage     int        `json:"actualPage"`
	ItemsPerPage   int        `json:"itemsPerPage"`
	NumPaginations int        `json:"numPaginations"`
	HiddenResults  bool       `json:"hiddenResults"`
	Summary        []string   `json:"summary"`
	Filter         struct {
		LocationName string `json:"locationName"`
	} `json:"filter"`
	LowerRangePosition int  `json:"lowerRangePosition"`
	UpperRangePosition int  `json:"upperRangePosition"`
	Paginable          bool `json:"paginable"`
}

type Property struct {
	PropertyCode      string     `json:"propertyCode" gorm:"column:id"`
	Thumbnail         string     `json:"thumbnail"`
	ExternalReference string     `json:"externalReference"`
	NumPhotos         int        `json:"numPhotos"`
	Floor             string     `json:"floor"`
	Price             float64    `json:"price"`
	PropertyType      string     `json:"propertyType"`
	Operation         string     `json:"operation"`
	Size              float64    `json:"size"`
	Exterior          bool       `json:"exterior"`
	Rooms             int        `json:"rooms"`
	Bathrooms         int        `json:"bathrooms"`
	Address           string     `json:"address"`
	Province          string     `json:"province"`
	Municipality      string     `json:"municipality"`
	District          string     `json:"district"`
	Country           string     `json:"country"`
	Neighborhood      string     `json:"neighborhood"`
	Latitude          float64    `json:"latitude"`
	Longitude         float64    `json:"longitude"`
	ShowAddress       bool       `json:"showAddress"`
	URL               string     `json:"url"`
	HasVideo          bool       `json:"hasVideo"`
	Status            string     `json:"status"`
	NewDevelopment    bool       `json:"newDevelopment"`
	PriceDropValue    int        `json:"priceDropValue"`
	DropDate          int64      `json:"dropDate"`
	Favourite         bool       `json:"favourite"`
	NewProperty       bool       `json:"newProperty"`
	Multimedia        Multimedia `json:"multimedia" gorm:"-"`
	ContactInfo       struct {
		CommercialName string `json:"commercialName"`
		Phone1         struct {
			PhoneNumber                 string `json:"phoneNumber"`
			FormattedPhone              string `json:"formattedPhone"`
			PhoneNumberForMobileDialing string `json:"phoneNumberForMobileDialing"`
			NationalNumber              bool   `json:"nationalNumber"`
		} `json:"phone1"`
		ContactName        string `json:"contactName"`
		UserType           string `json:"userType"`
		ContactMethod      string `json:"contactMethod"`
		MicrositeShortName string `json:"micrositeShortName"`
		Professional       bool   `json:"professional"`
	} `json:"contactInfo" gorm:"-"`
	HasLift             bool    `json:"hasLift"`
	PriceDropPercentage int     `json:"priceDropPercentage"`
	PriceByArea         float64 `json:"priceByArea"`
	Features            struct {
		HasSwimmingPool    bool `json:"hasSwimmingPool"`
		HasTerrace         bool `json:"hasTerrace"`
		HasAirConditioning bool `json:"hasAirConditioning"`
		HasBoxRoom         bool `json:"hasBoxRoom"`
		HasGarden          bool `json:"hasGarden"`
	} `json:"features" gorm:"-"`
	DetailedType struct {
		Typology string `json:"typology"`
	} `json:"detailedType" gorm:"-"`
	SuggestedTexts struct {
		Subtitle string `json:"subtitle"`
		Title    string `json:"title"`
	} `json:"suggestedTexts" gorm:"-"`
	HasPlan               bool `json:"hasPlan"`
	Has3DTour             bool `json:"has3DTour"`
	Has360                bool `json:"has360"`
	HasStaging            bool `json:"hasStaging"`
	TopNewDevelopment     bool `json:"topNewDevelopment"`
	PreferenceHighlight   bool `json:"preferenceHighlight"`
	TopHighlight          bool `json:"topHighlight"`
	UrgentVisualHighlight bool `json:"urgentVisualHighlight"`
	VisualHighlight       bool `json:"visualHighlight"`
}

type Multimedia struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

type PropertyDetails struct {
	Adid                 int     `json:"adid"`
	Price                float64 `json:"price"`
	Operation            string  `json:"operation"`
	PropertyType         string  `json:"propertyType"`
	ExtendedPropertyType string  `json:"extendedPropertyType"`
	HomeType             string  `json:"homeType"`
	State                string  `json:"state"`
	Multimedia           struct {
		Images []struct {
			URL           string `json:"url"`
			Tag           string `json:"tag"`
			LocalizedName string `json:"localizedName"`
		} `json:"images"`
		Videos            []interface{} `json:"videos"`
		VirtualTourImages []interface{} `json:"virtualTourImages"`
	} `json:"multimedia"`
	PropertyComment  string `json:"propertyComment"`
	HighlightComment string `json:"highlightComment"`
	Ubication        struct {
		Title                    string  `json:"title"`
		Latitude                 float64 `json:"latitude"`
		Longitude                float64 `json:"longitude"`
		HasHiddenAddress         bool    `json:"hasHiddenAddress"`
		AdministrativeAreaLevel4 string  `json:"administrativeAreaLevel4"`
		AdministrativeAreaLevel3 string  `json:"administrativeAreaLevel3"`
		AdministrativeAreaLevel2 string  `json:"administrativeAreaLevel2"`
		AdministrativeAreaLevel1 string  `json:"administrativeAreaLevel1"`
		LocationID               string  `json:"locationId"`
	} `json:"ubication"`
	Country     string `json:"country"`
	ContactInfo struct {
		CommercialName string `json:"commercialName"`
		Phone1         struct {
			PhoneNumber                 string `json:"phoneNumber"`
			FormattedPhone              string `json:"formattedPhone"`
			PhoneNumberForMobileDialing string `json:"phoneNumberForMobileDialing"`
			NationalNumber              bool   `json:"nationalNumber"`
		} `json:"phone1"`
		ContactName        string `json:"contactName"`
		UserType           string `json:"userType"`
		AgencyLogo         string `json:"agencyLogo"`
		ContactMethod      string `json:"contactMethod"`
		MicrositeShortName string `json:"micrositeShortName"`
		Address            struct {
			StreetName   string `json:"streetName"`
			StreetNumber int    `json:"streetNumber"`
			LocationName string `json:"locationName"`
			PostalCode   string `json:"postalCode"`
		} `json:"address"`
		AgentInfo struct {
			Name     string `json:"name"`
			ProAgent bool   `json:"proAgent"`
		} `json:"agentInfo"`
		InVirtualMicrosite bool `json:"inVirtualMicrosite"`
		Professional       bool `json:"professional"`
	} `json:"contactInfo"`
	MoreCharacteristics struct {
		ModificationDate        int64  `json:"modificationDate"`
		RoomNumber              int    `json:"roomNumber"`
		BathNumber              int    `json:"bathNumber"`
		ConstructedArea         int    `json:"constructedArea"`
		HousingFurnitures       string `json:"housingFurnitures"`
		Lift                    bool   `json:"lift"`
		AgencyIsABank           bool   `json:"agencyIsABank"`
		EnergyCertificationType string `json:"energyCertificationType"`
		FlatLocation            string `json:"flatLocation"`
		Status                  string `json:"status"`
	} `json:"moreCharacteristics"`
	TranslatedTexts struct {
		FloorNumberDescription      string `json:"floorNumberDescription"`
		CharacteristicsDescriptions []struct {
			Key     string   `json:"key"`
			Title   string   `json:"title"`
			Phrases []string `json:"phrases"`
		} `json:"characteristicsDescriptions"`
	} `json:"translatedTexts"`
	SuggestedTexts struct {
		Title string `json:"title"`
	} `json:"suggestedTexts"`
	DetailedType struct {
		Typology string `json:"typology"`
	} `json:"detailedType"`
	Comments []struct {
		PropertyComment  string `json:"propertyComment"`
		AutoTranslated   bool   `json:"autoTranslated"`
		Language         string `json:"language"`
		DefaultLanguage  bool   `json:"defaultLanguage"`
		HighlightComment string `json:"highlightComment,omitempty"`
	} `json:"comments"`
	DetailWebLink       string `json:"detailWebLink"`
	EnergyCertification struct {
		Prefix                  string `json:"prefix"`
		EnergyCertificationType string `json:"energyCertificationType"`
		Suffix                  string `json:"suffix"`
		HasIcon                 bool   `json:"hasIcon"`
	} `json:"energyCertification"`
	AllowsCounterOffers bool `json:"allowsCounterOffers"`
	PriceReferenceIndex struct {
		Mandatory bool `json:"mandatory"`
	} `json:"priceReferenceIndex"`
	Tracking struct {
		CommercialDataID int `json:"commercialDataId"`
	} `json:"tracking"`
}

func (user *Property) TableName() string {
	return "idealista_flats"
}
