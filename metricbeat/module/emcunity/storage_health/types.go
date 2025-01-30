package storage_health

import "time"

type Serial string

type ErrorResponse struct {
	Error struct {
		ErrorCode      int `json:"errorCode"`
		HTTPStatusCode int `json:"httpStatusCode"`
		Messages       []struct {
			EnUS string `json:"en-US"`
		} `json:"messages"`
		Created time.Time `json:"created"`
	} `json:"error"`
}

type UnityData struct {
	system           System_JSON
	pool             Pool_JSON
	poolUnit         PoolUnit_JSON
	lun              Lun_JSON
	storageProcesser StorageProcessor_JSON
	storageResource  StorageResource_JSON
	storageTier      StorageTier_JSON
	license          License_JSON
	ethernetPort     BasicEMCUnity_JSON
	fileInterface    BasicEMCUnity_JSON
	remoteSystem     BasicEMCUnity_JSON
	disk             Disk_JSON
	datastore        Datastore_JSON
	filesystem       FileSystem_JSON
	snap             Snap_JSON
	sasPort          SasPort_JSON
	powerSupply      BasicEMCUnity_JSON
	fan              BasicEMCUnity_JSON
	dae              Dae_JSON
	memoryModule     BasicEMCUnity_JSON
	battery          BasicEMCUnity_JSON
	ssd              BasicEMCUnity_JSON
	raidGroup        BasicEMCUnity_JSON
	treeQuota        TreeQuota_JSON
	diskGroup        DiskGroup_JSON
	cifsServer       BasicEMCUnity_JSON
	fastCache        FastCache_JSON
	fastVP           FastVP_JSON
	fcPort           BasicEMCUnity_JSON
	hostContainer    BasicEMCUnity_JSON
	hostInitiator    BasicEMCUnity_JSON
	host             BasicEMCUnity_JSON
	ioModule         BasicEMCUnity_JSON
	lcc              BasicEMCUnity_JSON
	nasServer        BasicEMCUnity_JSON
}

// These are most common data streams.

// This presently covers three scenarios EthernetPort, FileInterface, RemoteSystem
type BasicEMCUnity_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID              string `json:"id"`
			Name            string `json:"name"`
			NeedsReplacment bool   `json:"needsReplacement"`
			Size            int    `json:"size"`
			SizeTotal       int64  `json:"sizeTotal"`
			Health          struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
		} `json:"content"`
	} `json:"entries"`
}

// type BasicEMCUnityWithReplacement_JSON struct {
// 	// Base    string    `json:"@base"`
// 	// Updated time.Time `json:"updated"`
// 	// Links   []struct {
// 	// 	Rel  string `json:"rel"`
// 	// 	Href string `json:"href"`
// 	// } `json:"links"`
// 	Entries []struct {
// 		Content struct {
// 			ID              string `json:"id"`
// 			Name            string `json:"name"`
// 			NeedsReplacment bool   `json:"needsReplacement"`
// 			Health          struct {
// 				Value          int      `json:"value"`
// 				DescriptionIds []string `json:"descriptionIds"`
// 				Descriptions   []string `json:"descriptions"`
// 			} `json:"health"`
// 		} `json:"content"`
// 	} `json:"entries"`
// }

type System_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Name          string `json:"name"`
			Model         string `json:"model"`
			SerialNumber  string `json:"serialNumber"`
			InternalModel string `json:"internalModel"`
			Platform      string `json:"platform"`
			MacAddress    string `json:"macAddress"`
			// SystemUUID    string `json:"systemUUID"`
		} `json:"content"`
	} `json:"entries"`
}

type Pool_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Name                   string `json:"name"`
			Description            string `json:"description"`
			SizeFree               int64  `json:"sizeFree"`
			SizeTotal              int64  `json:"sizeTotal"`
			HarvestState           int    `json:"harvestState"`
			MetadataSizeSubscribed int64  `json:"metadataSizeSubscribed"`
			MetadataSizeUsed       int64  `json:"metadataSizeUsed"`
			RebalanceProgress      int    `json:"rebalanceProgress"`
			SizeSubscribed         int64  `json:"sizeSubscribed"`
			SizeUsed               int64  `json:"sizeUsed"`
			SnapSizeSubscribed     int64  `json:"snapSizeSubscribed"`
			SnapSizeUsed           int64  `json:"snapSizeUsed"`
		} `json:"content"`
	} `json:"entries"`
}

type PoolUnit_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Name      string `json:"name"`
			SizeTotal int64  `json:"sizeTotal"`
		} `json:"content"`
	} `json:"entries"`
}

type Lun_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			SizeTotal             int64 `json:"sizeTotal"`
			SizeAllocated         int   `json:"sizeAllocated"`
			MetadataSize          int64 `json:"metadataSize"`
			MetadataSizeAllocated int64 `json:"metadataSizeAllocated"`
			SnapsSize             int   `json:"snapsSize"`
			SnapsSizeAllocated    int   `json:"snapsSizeAllocated"`
		} `json:"content"`
	} `json:"entries"`
}
type StorageProcessor_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Model string `json:"model"`
			Name  string `json:"name"`
		} `json:"content"`
	} `json:"entries"`
}

type StorageResource_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			SizeTotal             int64 `json:"sizeTotal"`
			SizeAllocated         int   `json:"sizeAllocated"`
			MetadataSize          int64 `json:"metadataSize"`
			MetadataSizeAllocated int64 `json:"metadataSizeAllocated"`
			SnapsSizeTotal        int   `json:"snapsSizeTotal"`
			SnapsSizeAllocated    int   `json:"snapsSizeAllocated"`
			SnapCount             int   `json:"snapCount"`
		} `json:"content"`
	} `json:"entries"`
}

type StorageTier_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID                 string `json:"id"`
			DisksTotal         int    `json:"disksTotal"`
			DisksUnused        int    `json:"disksUnused"`
			VirtualDisksTotal  int    `json:"virtualDisksTotal"`
			VirtualDisksUnused int    `json:"virtualDisksUnused"`
			SizeTotal          int64  `json:"sizeTotal"`
			SizeFree           int64  `json:"sizeFree"`
		} `json:"content"`
	} `json:"entries"`
}

type License_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID          string    `json:"id"`
			Name        string    `json:"name"`
			IsInstalled bool      `json:"isInstalled"`
			IsValid     bool      `json:"isValid"`
			IsPermanent bool      `json:"isPermanent"`
			Expires     time.Time `json:"expires"`
			Feature     struct {
				ID string `json:"id"`
			} `json:"feature"`
		} `json:"content"`
	} `json:"entries"`
}

// const Disk_API = "/api/types/datastore/instances?fields=health,size,rawSize,vendorSize&compact=true"
// https://docs.vmware.com/en/Management-Packs-for-vRealize-Operations/4.0/dell-emc-unity/GUID-C6A6B0A5-B9E2-4061-92AC-6169FC546CC6.html
type Disk_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Name       string `json:"name"`
			Size       int64  `json:"size"`
			RawSize    int64  `json:"rawSize"`
			VendorSize int64  `json:"vendorSize"`
		} `json:"content"`
	} `json:"entries"`
}

// const DataStore_API = "/api/types/disk/instances?fields=sizeTotal,sizeUsed&compact=true"

type Datastore_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			SizeTotal int64  `json:"sizeTotal"`
			SizeUsed  int64  `json:"sizeUsed"`
		} `json:"content"`
	} `json:"entries"`
}

// const Disk_API = "/api/types/filesystem/instances?fields=name,health,metadataSize,metadataSizeAllocated,sizeAllocated,sizeTotal,sizeUsed,snapCount,snapsSize,snapsSizeAllocated&compact=true"
type FileSystem_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			MetadataSize          int64 `json:"metadataSize"`
			MetadataSizeAllocated int64 `json:"metadataSizeAllocated"`
			SizeTotal             int64 `json:"sizeTotal"`
			SizeUsed              int64 `json:"sizeUsed"`
			SizeAllocated         int64 `json:"sizeAllocated"`
			SnapsSize             int64 `json:"snapsSize"`
			SnapsSizeAllocated    int64 `json:"snapsSizeAllocated"`
			SnapCount             int   `json:"snapCount"`
		} `json:"content"`
	} `json:"entries"`
}

// const Snap_API = "/api/types/snap/instances?fields=name,size,state,expirationTime,creationTime&compact=true"
type Snap_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID             string    `json:"id"`
			Name           string    `json:"name"`
			Size           int64     `json:"size"`
			State          int       `json:"state"`
			ExpirationTime time.Time `json:"expirationTime"`
			CreationTime   time.Time `json:"creationTime"`
		} `json:"content"`
	} `json:"entries"`
}

// const SasPort_API = "/api/types/sasPort/instances?fields=name,needsReplacement,health&compact=true"
type SasPort_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID              string `json:"id"`
			Name            string `json:"name"`
			NeedsReplacment bool   `json:"needsReplacement"`
			Health          struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
		} `json:"content"`
	} `json:"entries"`
}

// const dae_API = "/api/types/dae/instances?fields=health,currentPower,avgPower,maxPower,currentTemperature,avgTemperature,maxTemperature&compact=true"
type Dae_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		// Base    string    `json:"@base"`
		// Updated time.Time `json:"updated"`
		// Links   []struct {
		// 	Rel  string `json:"rel"`
		// 	Href string `json:"href"`
		// } `json:"links"`
		Content struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			Model              string `json:"model"`
			NeedsReplacment    bool   `json:"needsReplacement"`
			CurrentPower       int    `json:"currentPower"`
			AvgPower           int    `json:"avgPower"`
			MaxPower           int    `json:"maxPower"`
			CurrentTemperature int    `json:"currentTemperature"`
			AvgTemperature     int    `json:"avgTemperature"`
			MaxTemperature     int    `json:"maxTemperature"`
			Health             struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
		} `json:"content"`
	} `json:"entries"`
}

type TreeQuota_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID        string `json:"id"`
			SoftLimit int64  `json:"softLimit"`
			HardLimit int64  `json:"hardLimit"`
			SizeUsed  int64  `json:"sizeUsed"`
			State     int    `json:"state"`
		} `json:"content"`
	} `json:"entries"`
}

// const DiskGroup_API = "/api/types/diskGroup/instances?fields=name,advertisedSize,diskSize,hotSparePolicyStatus,minHotSpareCandidates,rpm,speed,totalDisks,unconfiguredDisks&compact=true"
type DiskGroup_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID                    string `json:"id"`
			Name                  string `json:"name"`
			AdvertisedSize        int64  `json:"advertisedSize"`
			DiskSize              int64  `json:"diskSize"`
			HotSparePolicyStatus  int    `json:"hotSparePolicyStatus"`
			MinHotSpareCandidates int    `json:"minHotSpareCandidates"`
			Rpm                   int    `json:"rpm"`
			Speed                 int64  `json:"speed"`
			TotalDisks            int    `json:"totalDisks"`
			UnconfiguredDisks     int    `json:"unconfiguredDisks"`
		} `json:"content"`
	} `json:"entries"`
}

type FastCache_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID     string `json:"id"`
			Health struct {
				Value          int      `json:"value"`
				DescriptionIds []string `json:"descriptionIds"`
				Descriptions   []string `json:"descriptions"`
			} `json:"health"`
			Name          string `json:"name"`
			SizeFree      int64  `json:"sizeFree"`
			SizeTotal     int64  `json:"sizeTotal"`
			NumberOfDisks int    `json:"numberOfDisks"`
		} `json:"content"`
	} `json:"entries"`
}

// const FastVP_API = "/api/types/fastVP/instances?fields=id,isScheduleEnabled,relocationRate,sizeMovingUp,sizeMovingDown,sizeMovingWithin,status,relocationDurationEstimate&compact=true"
type FastVP_JSON struct {
	// Base    string    `json:"@base"`
	// Updated time.Time `json:"updated"`
	// Links   []struct {
	// 	Rel  string `json:"rel"`
	// 	Href string `json:"href"`
	// } `json:"links"`
	Entries []struct {
		Content struct {
			ID                         string `json:"id"`
			IsScheduleEnabled          bool   `json:"isScheduleEnabled"`
			RelocationRate             int    `json:"relocationRate"`
			SizeMovingUp               int64  `json:"sizeMovingUp"`
			SizeMovingDown             int64  `json:"sizeMovingDown"`
			SizeMovingWithin           int64  `json:"sizeMovingWithin"`
			Status                     int    `json:"status"`
			RelocationDurationEstimate string `json:"relocationDurationEstimate"`
		} `json:"content"`
	} `json:"entries"`
}
