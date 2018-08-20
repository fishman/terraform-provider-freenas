package freenas

import (
	"log"

	"github.com/hashicorp/terraform/terraform"
)

// resourceFreenasNfsShareMigrateState is the master state migration function for
// the freenas_nfs_share resource.
func resourceFreenasNfsShareMigrateState(version int, os *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	// Guard against a nil state.
	if os == nil {
		return nil, nil
	}

	// Guard against empty state, can't do anything with it
	if os.Empty() {
		return os, nil
	}

	var migrateFunc func(*terraform.InstanceState, interface{}) error
	switch version {
	case 0:
		log.Printf("[DEBUG] Migrating freenas_nfs_share state: old v%d state: %#v", version, os)
		migrateFunc = resourceFreenasNfsShareMigrateStateV1
	default:
		// Migration is complete
		log.Printf("[DEBUG] Migrating freenas_nfs_share state: completed v%d state: %#v", version, os)
		return os, nil
	}
	if err := migrateFunc(os, meta); err != nil {
		return nil, err
	}
	version++
	log.Printf("[DEBUG] Migrating freenas_nfs_share state: new v%d state: %#v", version, os)
	return resourceFreenasNfsShareMigrateState(version, os, meta)
}

// resourceFreenasNfsShareMigrateStateV1 migrates the state of the freenas_nfs_share
// from version 0 to version 1.
func resourceFreenasNfsShareMigrateStateV1(s *terraform.InstanceState, meta interface{}) error {
	// Our path for migration here is pretty much the same as our import path, so
	// we just leverage that functionality.
	//
	// We just need the path and the datacenter to proceed. We don't have an
	// analog in for existing_path in the new resource, so we just drop that on
	// the floor.

	return nil
}
