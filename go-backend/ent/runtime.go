// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/avored/go-ecommerce/ent/adminuser"
	"github.com/avored/go-ecommerce/ent/category"
	"github.com/avored/go-ecommerce/ent/role"
	"github.com/avored/go-ecommerce/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminuserMixin := schema.AdminUser{}.Mixin()
	adminuserMixinFields0 := adminuserMixin[0].Fields()
	_ = adminuserMixinFields0
	adminuserFields := schema.AdminUser{}.Fields()
	_ = adminuserFields
	// adminuserDescCreatedAt is the schema descriptor for created_at field.
	adminuserDescCreatedAt := adminuserMixinFields0[0].Descriptor()
	// adminuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	adminuser.DefaultCreatedAt = adminuserDescCreatedAt.Default.(func() time.Time)
	// adminuserDescUpdatedAt is the schema descriptor for updated_at field.
	adminuserDescUpdatedAt := adminuserMixinFields0[1].Descriptor()
	// adminuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	adminuser.DefaultUpdatedAt = adminuserDescUpdatedAt.Default.(func() time.Time)
	// adminuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	adminuser.UpdateDefaultUpdatedAt = adminuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	categoryMixin := schema.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryMixinFields0[0].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryMixinFields0[1].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
}
