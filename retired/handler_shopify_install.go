package fastseer

/*func (s *Server) handleReInstallSearchForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		shop := params.Get("shop")
		shopClient, _ := shopify.ShopClientConfig(shop, s.ClientsStore)

		err := shopify.InstallSearchFormThemeAsset(shopClient)
		if err != nil {
			logger.Error(shop, err.Error())
			s.SetFlashMessage(w, r, err.Error())
		}

		// re-install script tag
		_, err = shopify.InstallShopScriptTag(shopClient, s.AppDomain(r))
		if err != nil {
			logger.Error(shop, err.Error())
			s.SetFlashMessage(w, r, err.Error())
			return
		}

		s.SetFlashMessage(w, r, "Reinstalled Theme Assets")

		redir := r.Referer()
		http.Redirect(w, r, redir, 307)
	}
}*/
