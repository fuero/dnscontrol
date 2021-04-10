package models

import "fmt"

// SoaInfo contains the parts of the default SOA settings.
type SoaInfo struct {
	Ns      string `json:"master"`
	Mbox    string `json:"mbox"`
	Serial  uint32 `json:"serial"`
	Refresh uint32 `json:"refresh"`
	Retry   uint32 `json:"retry"`
	Expire  uint32 `json:"expire"`
	Minttl  uint32 `json:"minttl"`
	TTL     uint32 `json:"ttl,omitempty"`
}

func (s SoaInfo) String() string {
	return fmt.Sprintf("%s %s %d %d %d %d %d %d", s.Ns, s.Mbox, s.Serial, s.Refresh, s.Retry, s.Expire, s.Minttl, s.TTL)
}

func MakeSoa(origin string, defSoa *SoaInfo, existing, desired *RecordConfig) (*RecordConfig, uint32) {
	// Create a SOA record.  Take data from desired, existing, default,
	// or hardcoded defaults.
	soaRec := RecordConfig{}
	soaRec.SetLabel("@", origin)

	if defSoa == nil {
		defSoa = &SoaInfo{}
	}
	if existing == nil {
		existing = &RecordConfig{}
	}

	if desired == nil {
		desired = &RecordConfig{}
	}

	soaRec.TTL = firstNonZero(desired.TTL, defSoa.TTL, existing.TTL, DefaultTTL)
	soaRec.SetTargetSOA(
		firstNonNull(desired.GetTargetField(), existing.GetTargetField(), defSoa.Ns, "DEFAULT_NOT_SET."),
		firstNonNull(desired.SoaMbox, existing.SoaMbox, defSoa.Mbox, "DEFAULT_NOT_SET."),
		firstNonZero(desired.SoaSerial, existing.SoaSerial, defSoa.Serial, 1),
		firstNonZero(desired.SoaRefresh, existing.SoaRefresh, defSoa.Refresh, 3600),
		firstNonZero(desired.SoaRetry, existing.SoaRetry, defSoa.Retry, 600),
		firstNonZero(desired.SoaExpire, existing.SoaExpire, defSoa.Expire, 604800),
		firstNonZero(desired.SoaMinttl, existing.SoaMinttl, defSoa.Minttl, 1440),
	)

	return &soaRec, generateSerial(soaRec.SoaSerial)
}

func firstNonNull(items ...string) string {
	for _, item := range items {
		if item != "" {
			return item
		}
	}
	return "FAIL"
}

func firstNonZero(items ...uint32) uint32 {
	for _, item := range items {
		if item != 0 {
			return item
		}
	}
	return 999
}
