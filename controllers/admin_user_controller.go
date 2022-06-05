package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"golang.org/x/crypto/bcrypt"
)

func AdminLoginGetHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{"templates/admin/auth/login.html"}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, nil)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

func AdminLoginPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)

	userModel, _ := services.NewAdminUserOps(r.Context()).AdminUserGetByEmail(r.FormValue("email"))

	match := CheckPasswordHash(r.FormValue("password"), userModel.Password)

	NewUrl := "/admin/login"

	fmt.Printf("Match: %v", userModel)
	// Set user as Authenticated
	if match {
		session.Values["AdminUser_Authenticated"] = true
		session.Values["AdminUser_FirstName"] = userModel.FirstName
		session.Values["AdminUser_LastName"] = userModel.LastName
		session.Values["AdminUser_Email"] = userModel.Email
		session.Save(r, w)
		NewUrl = "/admin"
	}

	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}

func HandleAdminUserProfile(w http.ResponseWriter, r *http.Request) {

	existingAdminUser, _ := services.NewAdminUserOps(r.Context()).AdminUserGetByEmail("admin@admin.com")

	if existingAdminUser == nil {
		var newAdminUser ent.AdminUser
		newAdminUser.Email = "admin@admin.com"
		newAdminUser.Password = "admin123"
		newAdminUser.FirstName = "admin"
		newAdminUser.LastName = "admin"

		pass, err := HashPassword(newAdminUser.Password)

		if err != nil {
			fmt.Printf("Error while hashing password: %s", err)
		}
		newAdminUser.Password = string(pass)

		fmt.Printf("Admin User STRUCT ***** !!! %v", newAdminUser)

		adminUser, err := services.NewAdminUserOps(r.Context()).AdminUserCreate(newAdminUser)
		if err != nil {
			http.Error(w, "Error while creating an admin user account", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Admin User Profile Create successfully !!! %v", adminUser)
	}

	avoredCategory, _ := services.NewCategoryOps(r.Context()).CategoryGetBySlug("avored")

	if avoredCategory == nil {
		var newAvoRedCategory ent.Category
		newAvoRedCategory.Name = "AvoRed"
		newAvoRedCategory.Slug = "avored"
		newAvoRedCategory.MetaTitle = "AvoRed an elegant GO E commerce"
		newAvoRedCategory.MetaDescription = "AvoRed an elegant GO E commerce"

		newAddedCategory, err := services.NewCategoryOps(r.Context()).CategoryCreate(newAvoRedCategory)
		if err != nil {
			http.Error(w, "Error while creating an category", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Admin User Profile Create successfully !!! %v", newAddedCategory)
	}
	goCategory, _ := services.NewCategoryOps(r.Context()).CategoryGetBySlug("go")

	if goCategory == nil {
		var newgoCategory ent.Category
		newgoCategory.Name = "GO"
		newgoCategory.Slug = "go"
		newgoCategory.MetaTitle = "GO meta title"
		newgoCategory.MetaDescription = "GO meta description"

		newAddedCategory, err := services.NewCategoryOps(r.Context()).CategoryCreate(newgoCategory)
		if err != nil {
			http.Error(w, "Error while creating an category", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Admin User Profile Create successfully !!! %v", newAddedCategory)
	}
	fiberCategory, _ := services.NewCategoryOps(r.Context()).CategoryGetBySlug("fiber")

	if fiberCategory == nil {
		var newFiberCategory ent.Category
		newFiberCategory.Name = "Fiber"
		newFiberCategory.Slug = "fiber"
		newFiberCategory.MetaTitle = "Fiber meta title"
		newFiberCategory.MetaDescription = "Fiber meta description"

		newAddedCategory, err := services.NewCategoryOps(r.Context()).CategoryCreate(newFiberCategory)
		if err != nil {
			http.Error(w, "Error while creating an category", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Admin User Profile Create successfully !!! %v", newAddedCategory)
	}

	revelCategory, _ := services.NewCategoryOps(r.Context()).CategoryGetBySlug("revel")

	if revelCategory == nil {
		var newRevelCategory ent.Category
		newRevelCategory.Name = "Fiber"
		newRevelCategory.Slug = "fiber"
		newRevelCategory.MetaTitle = "Fiber meta title"
		newRevelCategory.MetaDescription = "Fiber meta description"

		newAddedCategory, err := services.NewCategoryOps(r.Context()).CategoryCreate(newRevelCategory)
		if err != nil {
			http.Error(w, "Error while creating an category", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Admin User Profile Create successfully !!! %v", newAddedCategory)
	}

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
