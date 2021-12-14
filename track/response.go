package spotify

type GetUsersSavedTracksResponse struct {
	Href  string `json:"href"`
	Items []struct {
	} `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type GetTracksResponse struct {
	Album struct {
		AlbumType        string   `json:"album_type"`
		TotalTracks      int      `json:"total_tracks"`
		AvailableMarkets []string `json:"available_markets"`
		ExternalUrls     struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		Id     string `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		Restrictions         struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Type       string `json:"type"`
		Uri        string `json:"uri"`
		AlbumGroup string `json:"album_group"`
		Artists    []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"artists"`
	} `json:"album"`
	Artists []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"followers"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		Id     string   `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		Uri        string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int      `json:"disc_number"`
	DurationMs       int      `json:"duration_ms"`
	Explicit         bool     `json:"explicit"`
	ExternalIds      struct {
		Isrc string `json:"isrc"`
		Ean  string `json:"ean"`
		Upc  string `json:"upc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href       string `json:"href"`
	Id         string `json:"id"`
	IsPlayable bool   `json:"is_playable"`
	LinkedFrom struct {
		Album struct {
			AlbumType        string   `json:"album_type"`
			TotalTracks      int      `json:"total_tracks"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			Id     string `json:"id"`
			Images []struct {
				Url    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			Restrictions         struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Type       string `json:"type"`
			Uri        string `json:"uri"`
			AlbumGroup string `json:"album_group"`
			Artists    []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			Id     string   `json:"id"`
			Images []struct {
				Url    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			Uri        string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
			Ean  string `json:"ean"`
			Upc  string `json:"upc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href       string `json:"href"`
		Id         string `json:"id"`
		IsPlayable bool   `json:"is_playable"`
		LinkedFrom struct {
			Album struct {
				AlbumType        string   `json:"album_type"`
				TotalTracks      int      `json:"total_tracks"`
				AvailableMarkets []string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				Id     string `json:"id"`
				Images []struct {
					Url    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				Restrictions         struct {
					Reason string `json:"reason"`
				} `json:"restrictions"`
				Type       string `json:"type"`
				Uri        string `json:"uri"`
				AlbumGroup string `json:"album_group"`
				Artists    []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					Id   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Uri  string `json:"uri"`
				} `json:"artists"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Followers struct {
					Href  string `json:"href"`
					Total int    `json:"total"`
				} `json:"followers"`
				Genres []string `json:"genres"`
				Href   string   `json:"href"`
				Id     string   `json:"id"`
				Images []struct {
					Url    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name       string `json:"name"`
				Popularity int    `json:"popularity"`
				Type       string `json:"type"`
				Uri        string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalIds      struct {
				Isrc string `json:"isrc"`
				Ean  string `json:"ean"`
				Upc  string `json:"upc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href       string `json:"href"`
			Id         string `json:"id"`
			IsPlayable bool   `json:"is_playable"`
			LinkedFrom struct {
				Album struct {
					AlbumType        string   `json:"album_type"`
					TotalTracks      int      `json:"total_tracks"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalUrls     struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					Id     string `json:"id"`
					Images []struct {
						Url    string `json:"url"`
						Height int    `json:"height"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string `json:"name"`
					ReleaseDate          string `json:"release_date"`
					ReleaseDatePrecision string `json:"release_date_precision"`
					Restrictions         struct {
						Reason string `json:"reason"`
					} `json:"restrictions"`
					Type       string `json:"type"`
					Uri        string `json:"uri"`
					AlbumGroup string `json:"album_group"`
					Artists    []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						Id   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						Uri  string `json:"uri"`
					} `json:"artists"`
				} `json:"album"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Followers struct {
						Href  string `json:"href"`
						Total int    `json:"total"`
					} `json:"followers"`
					Genres []string `json:"genres"`
					Href   string   `json:"href"`
					Id     string   `json:"id"`
					Images []struct {
						Url    string `json:"url"`
						Height int    `json:"height"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name       string `json:"name"`
					Popularity int    `json:"popularity"`
					Type       string `json:"type"`
					Uri        string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMs       int      `json:"duration_ms"`
				Explicit         bool     `json:"explicit"`
				ExternalIds      struct {
					Isrc string `json:"isrc"`
					Ean  string `json:"ean"`
					Upc  string `json:"upc"`
				} `json:"external_ids"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href       string `json:"href"`
				Id         string `json:"id"`
				IsPlayable bool   `json:"is_playable"`
				LinkedFrom struct {
					Album struct {
						AlbumType        string   `json:"album_type"`
						TotalTracks      int      `json:"total_tracks"`
						AvailableMarkets []string `json:"available_markets"`
						ExternalUrls     struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href   string `json:"href"`
						Id     string `json:"id"`
						Images []struct {
							Url    string `json:"url"`
							Height int    `json:"height"`
							Width  int    `json:"width"`
						} `json:"images"`
						Name                 string `json:"name"`
						ReleaseDate          string `json:"release_date"`
						ReleaseDatePrecision string `json:"release_date_precision"`
						Restrictions         struct {
							Reason string `json:"reason"`
						} `json:"restrictions"`
						Type       string `json:"type"`
						Uri        string `json:"uri"`
						AlbumGroup string `json:"album_group"`
						Artists    []struct {
							ExternalUrls struct {
								Spotify string `json:"spotify"`
							} `json:"external_urls"`
							Href string `json:"href"`
							Id   string `json:"id"`
							Name string `json:"name"`
							Type string `json:"type"`
							Uri  string `json:"uri"`
						} `json:"artists"`
					} `json:"album"`
					Artists []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Followers struct {
							Href  string `json:"href"`
							Total int    `json:"total"`
						} `json:"followers"`
						Genres []string `json:"genres"`
						Href   string   `json:"href"`
						Id     string   `json:"id"`
						Images []struct {
							Url    string `json:"url"`
							Height int    `json:"height"`
							Width  int    `json:"width"`
						} `json:"images"`
						Name       string `json:"name"`
						Popularity int    `json:"popularity"`
						Type       string `json:"type"`
						Uri        string `json:"uri"`
					} `json:"artists"`
					AvailableMarkets []string `json:"available_markets"`
					DiscNumber       int      `json:"disc_number"`
					DurationMs       int      `json:"duration_ms"`
					Explicit         bool     `json:"explicit"`
					ExternalIds      struct {
						Isrc string `json:"isrc"`
						Ean  string `json:"ean"`
						Upc  string `json:"upc"`
					} `json:"external_ids"`
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href       string `json:"href"`
					Id         string `json:"id"`
					IsPlayable bool   `json:"is_playable"`
					LinkedFrom struct {
						Album struct {
							AlbumType        string   `json:"album_type"`
							TotalTracks      int      `json:"total_tracks"`
							AvailableMarkets []string `json:"available_markets"`
							ExternalUrls     struct {
								Spotify string `json:"spotify"`
							} `json:"external_urls"`
							Href   string `json:"href"`
							Id     string `json:"id"`
							Images []struct {
								Url    string `json:"url"`
								Height int    `json:"height"`
								Width  int    `json:"width"`
							} `json:"images"`
							Name                 string `json:"name"`
							ReleaseDate          string `json:"release_date"`
							ReleaseDatePrecision string `json:"release_date_precision"`
							Restrictions         struct {
								Reason string `json:"reason"`
							} `json:"restrictions"`
							Type       string `json:"type"`
							Uri        string `json:"uri"`
							AlbumGroup string `json:"album_group"`
							Artists    []struct {
								ExternalUrls struct {
									Spotify string `json:"spotify"`
								} `json:"external_urls"`
								Href string `json:"href"`
								Id   string `json:"id"`
								Name string `json:"name"`
								Type string `json:"type"`
								Uri  string `json:"uri"`
							} `json:"artists"`
						} `json:"album"`
						Artists []struct {
							ExternalUrls struct {
								Spotify string `json:"spotify"`
							} `json:"external_urls"`
							Followers struct {
								Href  string `json:"href"`
								Total int    `json:"total"`
							} `json:"followers"`
							Genres []string `json:"genres"`
							Href   string   `json:"href"`
							Id     string   `json:"id"`
							Images []struct {
								Url    string `json:"url"`
								Height int    `json:"height"`
								Width  int    `json:"width"`
							} `json:"images"`
							Name       string `json:"name"`
							Popularity int    `json:"popularity"`
							Type       string `json:"type"`
							Uri        string `json:"uri"`
						} `json:"artists"`
						AvailableMarkets []string `json:"available_markets"`
						DiscNumber       int      `json:"disc_number"`
						DurationMs       int      `json:"duration_ms"`
						Explicit         bool     `json:"explicit"`
						ExternalIds      struct {
							Isrc string `json:"isrc"`
							Ean  string `json:"ean"`
							Upc  string `json:"upc"`
						} `json:"external_ids"`
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href       string `json:"href"`
						Id         string `json:"id"`
						IsPlayable bool   `json:"is_playable"`
						LinkedFrom struct {
							Album struct {
								AlbumType        string   `json:"album_type"`
								TotalTracks      int      `json:"total_tracks"`
								AvailableMarkets []string `json:"available_markets"`
								ExternalUrls     struct {
									Spotify string `json:"spotify"`
								} `json:"external_urls"`
								Href   string `json:"href"`
								Id     string `json:"id"`
								Images []struct {
									Url    string `json:"url"`
									Height int    `json:"height"`
									Width  int    `json:"width"`
								} `json:"images"`
								Name                 string `json:"name"`
								ReleaseDate          string `json:"release_date"`
								ReleaseDatePrecision string `json:"release_date_precision"`
								Restrictions         struct {
									Reason string `json:"reason"`
								} `json:"restrictions"`
								Type       string `json:"type"`
								Uri        string `json:"uri"`
								AlbumGroup string `json:"album_group"`
								Artists    []struct {
									ExternalUrls struct {
										Spotify string `json:"spotify"`
									} `json:"external_urls"`
									Href string `json:"href"`
									Id   string `json:"id"`
									Name string `json:"name"`
									Type string `json:"type"`
									Uri  string `json:"uri"`
								} `json:"artists"`
							} `json:"album"`
							Artists []struct {
								ExternalUrls struct {
									Spotify string `json:"spotify"`
								} `json:"external_urls"`
								Followers struct {
									Href  string `json:"href"`
									Total int    `json:"total"`
								} `json:"followers"`
								Genres []string `json:"genres"`
								Href   string   `json:"href"`
								Id     string   `json:"id"`
								Images []struct {
									Url    string `json:"url"`
									Height int    `json:"height"`
									Width  int    `json:"width"`
								} `json:"images"`
								Name       string `json:"name"`
								Popularity int    `json:"popularity"`
								Type       string `json:"type"`
								Uri        string `json:"uri"`
							} `json:"artists"`
							AvailableMarkets []string `json:"available_markets"`
							DiscNumber       int      `json:"disc_number"`
							DurationMs       int      `json:"duration_ms"`
							Explicit         bool     `json:"explicit"`
							ExternalIds      struct {
								Isrc string `json:"isrc"`
								Ean  string `json:"ean"`
								Upc  string `json:"upc"`
							} `json:"external_ids"`
							ExternalUrls struct {
								Spotify string `json:"spotify"`
							} `json:"external_urls"`
							Href       string `json:"href"`
							Id         string `json:"id"`
							IsPlayable bool   `json:"is_playable"`
							LinkedFrom struct {
								Album struct {
									AlbumType        string   `json:"album_type"`
									TotalTracks      int      `json:"total_tracks"`
									AvailableMarkets []string `json:"available_markets"`
									ExternalUrls     struct {
										Spotify string `json:"spotify"`
									} `json:"external_urls"`
									Href   string `json:"href"`
									Id     string `json:"id"`
									Images []struct {
										Url    string `json:"url"`
										Height int    `json:"height"`
										Width  int    `json:"width"`
									} `json:"images"`
									Name                 string `json:"name"`
									ReleaseDate          string `json:"release_date"`
									ReleaseDatePrecision string `json:"release_date_precision"`
									Restrictions         struct {
										Reason string `json:"reason"`
									} `json:"restrictions"`
									Type       string `json:"type"`
									Uri        string `json:"uri"`
									AlbumGroup string `json:"album_group"`
									Artists    []struct {
										ExternalUrls struct {
											Spotify string `json:"spotify"`
										} `json:"external_urls"`
										Href string `json:"href"`
										Id   string `json:"id"`
										Name string `json:"name"`
										Type string `json:"type"`
										Uri  string `json:"uri"`
									} `json:"artists"`
								} `json:"album"`
								Artists []struct {
									ExternalUrls struct {
										Spotify string `json:"spotify"`
									} `json:"external_urls"`
									Followers struct {
										Href  string `json:"href"`
										Total int    `json:"total"`
									} `json:"followers"`
									Genres []string `json:"genres"`
									Href   string   `json:"href"`
									Id     string   `json:"id"`
									Images []struct {
										Url    string `json:"url"`
										Height int    `json:"height"`
										Width  int    `json:"width"`
									} `json:"images"`
									Name       string `json:"name"`
									Popularity int    `json:"popularity"`
									Type       string `json:"type"`
									Uri        string `json:"uri"`
								} `json:"artists"`
								AvailableMarkets []string `json:"available_markets"`
								DiscNumber       int      `json:"disc_number"`
								DurationMs       int      `json:"duration_ms"`
								Explicit         bool     `json:"explicit"`
								ExternalIds      struct {
									Isrc string `json:"isrc"`
									Ean  string `json:"ean"`
									Upc  string `json:"upc"`
								} `json:"external_ids"`
								ExternalUrls struct {
									Spotify string `json:"spotify"`
								} `json:"external_urls"`
								Href       string `json:"href"`
								Id         string `json:"id"`
								IsPlayable bool   `json:"is_playable"`
								LinkedFrom struct {
									Album struct {
										AlbumType        string   `json:"album_type"`
										TotalTracks      int      `json:"total_tracks"`
										AvailableMarkets []string `json:"available_markets"`
										ExternalUrls     struct {
											Spotify string `json:"spotify"`
										} `json:"external_urls"`
										Href   string `json:"href"`
										Id     string `json:"id"`
										Images []struct {
											Url    string `json:"url"`
											Height int    `json:"height"`
											Width  int    `json:"width"`
										} `json:"images"`
										Name                 string `json:"name"`
										ReleaseDate          string `json:"release_date"`
										ReleaseDatePrecision string `json:"release_date_precision"`
										Restrictions         struct {
											Reason string `json:"reason"`
										} `json:"restrictions"`
										Type       string `json:"type"`
										Uri        string `json:"uri"`
										AlbumGroup string `json:"album_group"`
										Artists    []struct {
											ExternalUrls struct {
												Spotify string `json:"spotify"`
											} `json:"external_urls"`
											Href string `json:"href"`
											Id   string `json:"id"`
											Name string `json:"name"`
											Type string `json:"type"`
											Uri  string `json:"uri"`
										} `json:"artists"`
									} `json:"album"`
									Artists []struct {
										ExternalUrls struct {
											Spotify string `json:"spotify"`
										} `json:"external_urls"`
										Followers struct {
											Href  string `json:"href"`
											Total int    `json:"total"`
										} `json:"followers"`
										Genres []string `json:"genres"`
										Href   string   `json:"href"`
										Id     string   `json:"id"`
										Images []struct {
											Url    string `json:"url"`
											Height int    `json:"height"`
											Width  int    `json:"width"`
										} `json:"images"`
										Name       string `json:"name"`
										Popularity int    `json:"popularity"`
										Type       string `json:"type"`
										Uri        string `json:"uri"`
									} `json:"artists"`
									AvailableMarkets []string `json:"available_markets"`
									DiscNumber       int      `json:"disc_number"`
									DurationMs       int      `json:"duration_ms"`
									Explicit         bool     `json:"explicit"`
									ExternalIds      struct {
										Isrc string `json:"isrc"`
										Ean  string `json:"ean"`
										Upc  string `json:"upc"`
									} `json:"external_ids"`
									ExternalUrls struct {
										Spotify string `json:"spotify"`
									} `json:"external_urls"`
									Href       string `json:"href"`
									Id         string `json:"id"`
									IsPlayable bool   `json:"is_playable"`
									LinkedFrom struct {
										Album struct {
											AlbumType        string   `json:"album_type"`
											TotalTracks      int      `json:"total_tracks"`
											AvailableMarkets []string `json:"available_markets"`
											ExternalUrls     struct {
												Spotify string `json:"spotify"`
											} `json:"external_urls"`
											Href   string `json:"href"`
											Id     string `json:"id"`
											Images []struct {
												Url    string `json:"url"`
												Height int    `json:"height"`
												Width  int    `json:"width"`
											} `json:"images"`
											Name                 string `json:"name"`
											ReleaseDate          string `json:"release_date"`
											ReleaseDatePrecision string `json:"release_date_precision"`
											Restrictions         struct {
												Reason string `json:"reason"`
											} `json:"restrictions"`
											Type       string `json:"type"`
											Uri        string `json:"uri"`
											AlbumGroup string `json:"album_group"`
											Artists    []struct {
												ExternalUrls struct {
													Spotify string `json:"spotify"`
												} `json:"external_urls"`
												Href string `json:"href"`
												Id   string `json:"id"`
												Name string `json:"name"`
												Type string `json:"type"`
												Uri  string `json:"uri"`
											} `json:"artists"`
										} `json:"album"`
										Artists []struct {
											ExternalUrls struct {
												Spotify string `json:"spotify"`
											} `json:"external_urls"`
											Followers struct {
												Href  string `json:"href"`
												Total int    `json:"total"`
											} `json:"followers"`
											Genres []string `json:"genres"`
											Href   string   `json:"href"`
											Id     string   `json:"id"`
											Images []struct {
												Url    string `json:"url"`
												Height int    `json:"height"`
												Width  int    `json:"width"`
											} `json:"images"`
											Name       string `json:"name"`
											Popularity int    `json:"popularity"`
											Type       string `json:"type"`
											Uri        string `json:"uri"`
										} `json:"artists"`
										AvailableMarkets []string `json:"available_markets"`
										DiscNumber       int      `json:"disc_number"`
										DurationMs       int      `json:"duration_ms"`
										Explicit         bool     `json:"explicit"`
										ExternalIds      struct {
											Isrc string `json:"isrc"`
											Ean  string `json:"ean"`
											Upc  string `json:"upc"`
										} `json:"external_ids"`
										ExternalUrls struct {
											Spotify string `json:"spotify"`
										} `json:"external_urls"`
										Href       string `json:"href"`
										Id         string `json:"id"`
										IsPlayable bool   `json:"is_playable"`
										LinkedFrom struct {
											Album struct {
												AlbumType        string   `json:"album_type"`
												TotalTracks      int      `json:"total_tracks"`
												AvailableMarkets []string `json:"available_markets"`
												ExternalUrls     struct {
													Spotify string `json:"spotify"`
												} `json:"external_urls"`
												Href   string `json:"href"`
												Id     string `json:"id"`
												Images []struct {
													Url    string `json:"url"`
													Height int    `json:"height"`
													Width  int    `json:"width"`
												} `json:"images"`
												Name                 string `json:"name"`
												ReleaseDate          string `json:"release_date"`
												ReleaseDatePrecision string `json:"release_date_precision"`
												Restrictions         struct {
													Reason string `json:"reason"`
												} `json:"restrictions"`
												Type       string `json:"type"`
												Uri        string `json:"uri"`
												AlbumGroup string `json:"album_group"`
												Artists    []struct {
													ExternalUrls struct {
													} `json:"external_urls"`
													Href string `json:"href"`
													Id   string `json:"id"`
													Name string `json:"name"`
													Type string `json:"type"`
													Uri  string `json:"uri"`
												} `json:"artists"`
											} `json:"album"`
											Artists []struct {
												ExternalUrls struct {
													Spotify string `json:"spotify"`
												} `json:"external_urls"`
												Followers struct {
													Href  string `json:"href"`
													Total int    `json:"total"`
												} `json:"followers"`
												Genres []string `json:"genres"`
												Href   string   `json:"href"`
												Id     string   `json:"id"`
												Images []struct {
												} `json:"images"`
												Name       string `json:"name"`
												Popularity int    `json:"popularity"`
												Type       string `json:"type"`
												Uri        string `json:"uri"`
											} `json:"artists"`
											AvailableMarkets []string `json:"available_markets"`
											DiscNumber       int      `json:"disc_number"`
											DurationMs       int      `json:"duration_ms"`
											Explicit         bool     `json:"explicit"`
											ExternalIds      struct {
												Isrc string `json:"isrc"`
												Ean  string `json:"ean"`
												Upc  string `json:"upc"`
											} `json:"external_ids"`
											ExternalUrls struct {
												Spotify string `json:"spotify"`
											} `json:"external_urls"`
											Href       string `json:"href"`
											Id         string `json:"id"`
											IsPlayable bool   `json:"is_playable"`
											LinkedFrom struct {
												Album struct {
													AlbumType        string   `json:"album_type"`
													TotalTracks      int      `json:"total_tracks"`
													AvailableMarkets []string `json:"available_markets"`
													ExternalUrls     struct {
														Spotify string `json:"spotify"`
													} `json:"external_urls"`
													Href   string `json:"href"`
													Id     string `json:"id"`
													Images []struct {
													} `json:"images"`
													Name                 string `json:"name"`
													ReleaseDate          string `json:"release_date"`
													ReleaseDatePrecision string `json:"release_date_precision"`
													Restrictions         struct {
														Reason string `json:"reason"`
													} `json:"restrictions"`
													Type       string `json:"type"`
													Uri        string `json:"uri"`
													AlbumGroup string `json:"album_group"`
													Artists    []struct {
													} `json:"artists"`
												} `json:"album"`
												Artists []struct {
													ExternalUrls struct {
													} `json:"external_urls"`
													Followers struct {
													} `json:"followers"`
													Genres []string `json:"genres"`
													Href   string   `json:"href"`
													Id     string   `json:"id"`
													Images []struct {
													} `json:"images"`
													Name       string `json:"name"`
													Popularity int    `json:"popularity"`
													Type       string `json:"type"`
													Uri        string `json:"uri"`
												} `json:"artists"`
												AvailableMarkets []string `json:"available_markets"`
												DiscNumber       int      `json:"disc_number"`
												DurationMs       int      `json:"duration_ms"`
												Explicit         bool     `json:"explicit"`
												ExternalIds      struct {
													Isrc string `json:"isrc"`
													Ean  string `json:"ean"`
													Upc  string `json:"upc"`
												} `json:"external_ids"`
												ExternalUrls struct {
													Spotify string `json:"spotify"`
												} `json:"external_urls"`
												Href       string `json:"href"`
												Id         string `json:"id"`
												IsPlayable bool   `json:"is_playable"`
												LinkedFrom struct {
													Album struct {
														AlbumType        string   `json:"album_type"`
														TotalTracks      int      `json:"total_tracks"`
														AvailableMarkets []string `json:"available_markets"`
														ExternalUrls     struct {
														} `json:"external_urls"`
														Href   string `json:"href"`
														Id     string `json:"id"`
														Images []struct {
														} `json:"images"`
														Name                 string `json:"name"`
														ReleaseDate          string `json:"release_date"`
														ReleaseDatePrecision string `json:"release_date_precision"`
														Restrictions         struct {
														} `json:"restrictions"`
														Type       string `json:"type"`
														Uri        string `json:"uri"`
														AlbumGroup string `json:"album_group"`
														Artists    []struct {
														} `json:"artists"`
													} `json:"album"`
													Artists []struct {
														Genres []interface{} `json:"genres"`
														Images []interface{} `json:"images"`
													} `json:"artists"`
													AvailableMarkets []string `json:"available_markets"`
													DiscNumber       int      `json:"disc_number"`
													DurationMs       int      `json:"duration_ms"`
													Explicit         bool     `json:"explicit"`
													ExternalIds      struct {
														Isrc string `json:"isrc"`
														Ean  string `json:"ean"`
														Upc  string `json:"upc"`
													} `json:"external_ids"`
													ExternalUrls struct {
														Spotify string `json:"spotify"`
													} `json:"external_urls"`
													Href       string `json:"href"`
													Id         string `json:"id"`
													IsPlayable bool   `json:"is_playable"`
													LinkedFrom struct {
														Album struct {
															AvailableMarkets []interface{} `json:"available_markets"`
															Images           []interface{} `json:"images"`
															Artists          []interface{} `json:"artists"`
														} `json:"album"`
														Artists []struct {
														} `json:"artists"`
														AvailableMarkets []interface{} `json:"available_markets"`
														DiscNumber       int           `json:"disc_number"`
														DurationMs       int           `json:"duration_ms"`
														Explicit         bool          `json:"explicit"`
														ExternalIds      struct {
														} `json:"external_ids"`
														ExternalUrls struct {
														} `json:"external_urls"`
														Href       string `json:"href"`
														Id         string `json:"id"`
														IsPlayable bool   `json:"is_playable"`
														LinkedFrom struct {
															Artists          []interface{} `json:"artists"`
															AvailableMarkets []interface{} `json:"available_markets"`
														} `json:"linked_from"`
														Restrictions struct {
														} `json:"restrictions"`
														Name        string `json:"name"`
														Popularity  int    `json:"popularity"`
														PreviewUrl  string `json:"preview_url"`
														TrackNumber int    `json:"track_number"`
														Type        string `json:"type"`
														Uri         string `json:"uri"`
														IsLocal     bool   `json:"is_local"`
													} `json:"linked_from"`
													Restrictions struct {
														Reason string `json:"reason"`
													} `json:"restrictions"`
													Name        string `json:"name"`
													Popularity  int    `json:"popularity"`
													PreviewUrl  string `json:"preview_url"`
													TrackNumber int    `json:"track_number"`
													Type        string `json:"type"`
													Uri         string `json:"uri"`
													IsLocal     bool   `json:"is_local"`
												} `json:"linked_from"`
												Restrictions struct {
													Reason string `json:"reason"`
												} `json:"restrictions"`
												Name        string `json:"name"`
												Popularity  int    `json:"popularity"`
												PreviewUrl  string `json:"preview_url"`
												TrackNumber int    `json:"track_number"`
												Type        string `json:"type"`
												Uri         string `json:"uri"`
												IsLocal     bool   `json:"is_local"`
											} `json:"linked_from"`
											Restrictions struct {
												Reason string `json:"reason"`
											} `json:"restrictions"`
											Name        string `json:"name"`
											Popularity  int    `json:"popularity"`
											PreviewUrl  string `json:"preview_url"`
											TrackNumber int    `json:"track_number"`
											Type        string `json:"type"`
											Uri         string `json:"uri"`
											IsLocal     bool   `json:"is_local"`
										} `json:"linked_from"`
										Restrictions struct {
											Reason string `json:"reason"`
										} `json:"restrictions"`
										Name        string `json:"name"`
										Popularity  int    `json:"popularity"`
										PreviewUrl  string `json:"preview_url"`
										TrackNumber int    `json:"track_number"`
										Type        string `json:"type"`
										Uri         string `json:"uri"`
										IsLocal     bool   `json:"is_local"`
									} `json:"linked_from"`
									Restrictions struct {
										Reason string `json:"reason"`
									} `json:"restrictions"`
									Name        string `json:"name"`
									Popularity  int    `json:"popularity"`
									PreviewUrl  string `json:"preview_url"`
									TrackNumber int    `json:"track_number"`
									Type        string `json:"type"`
									Uri         string `json:"uri"`
									IsLocal     bool   `json:"is_local"`
								} `json:"linked_from"`
								Restrictions struct {
									Reason string `json:"reason"`
								} `json:"restrictions"`
								Name        string `json:"name"`
								Popularity  int    `json:"popularity"`
								PreviewUrl  string `json:"preview_url"`
								TrackNumber int    `json:"track_number"`
								Type        string `json:"type"`
								Uri         string `json:"uri"`
								IsLocal     bool   `json:"is_local"`
							} `json:"linked_from"`
							Restrictions struct {
								Reason string `json:"reason"`
							} `json:"restrictions"`
							Name        string `json:"name"`
							Popularity  int    `json:"popularity"`
							PreviewUrl  string `json:"preview_url"`
							TrackNumber int    `json:"track_number"`
							Type        string `json:"type"`
							Uri         string `json:"uri"`
							IsLocal     bool   `json:"is_local"`
						} `json:"linked_from"`
						Restrictions struct {
							Reason string `json:"reason"`
						} `json:"restrictions"`
						Name        string `json:"name"`
						Popularity  int    `json:"popularity"`
						PreviewUrl  string `json:"preview_url"`
						TrackNumber int    `json:"track_number"`
						Type        string `json:"type"`
						Uri         string `json:"uri"`
						IsLocal     bool   `json:"is_local"`
					} `json:"linked_from"`
					Restrictions struct {
						Reason string `json:"reason"`
					} `json:"restrictions"`
					Name        string `json:"name"`
					Popularity  int    `json:"popularity"`
					PreviewUrl  string `json:"preview_url"`
					TrackNumber int    `json:"track_number"`
					Type        string `json:"type"`
					Uri         string `json:"uri"`
					IsLocal     bool   `json:"is_local"`
				} `json:"linked_from"`
				Restrictions struct {
					Reason string `json:"reason"`
				} `json:"restrictions"`
				Name        string `json:"name"`
				Popularity  int    `json:"popularity"`
				PreviewUrl  string `json:"preview_url"`
				TrackNumber int    `json:"track_number"`
				Type        string `json:"type"`
				Uri         string `json:"uri"`
				IsLocal     bool   `json:"is_local"`
			} `json:"linked_from"`
			Restrictions struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewUrl  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			Uri         string `json:"uri"`
			IsLocal     bool   `json:"is_local"`
		} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewUrl  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		Uri         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	} `json:"linked_from"`
	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Name        string `json:"name"`
	Popularity  int    `json:"popularity"`
	PreviewUrl  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	Uri         string `json:"uri"`
	IsLocal     bool   `json:"is_local"`
}

type GetTracksAudioAnalysisResponse struct {
	Meta struct {
		AnalyzerVersion string  `json:"analyzer_version"`
		Platform        string  `json:"platform"`
		DetailedStatus  string  `json:"detailed_status"`
		StatusCode      int     `json:"status_code"`
		Timestamp       int     `json:"timestamp"`
		AnalysisTime    float64 `json:"analysis_time"`
		InputProcess    string  `json:"input_process"`
	} `json:"meta"`
	Track struct {
		NumSamples              int     `json:"num_samples"`
		Duration                float64 `json:"duration"`
		SampleMd5               string  `json:"sample_md5"`
		OffsetSeconds           int     `json:"offset_seconds"`
		WindowSeconds           int     `json:"window_seconds"`
		AnalysisSampleRate      int     `json:"analysis_sample_rate"`
		AnalysisChannels        int     `json:"analysis_channels"`
		EndOfFadeIn             int     `json:"end_of_fade_in"`
		StartOfFadeOut          float64 `json:"start_of_fade_out"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		Codestring              string  `json:"codestring"`
		CodeVersion             float64 `json:"code_version"`
		Echoprintstring         string  `json:"echoprintstring"`
		EchoprintVersion        float64 `json:"echoprint_version"`
		Synchstring             string  `json:"synchstring"`
		SynchVersion            int     `json:"synch_version"`
		Rhythmstring            string  `json:"rhythmstring"`
		RhythmVersion           int     `json:"rhythm_version"`
	} `json:"track"`
	Bars []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"bars"`
	Beats []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"beats"`
	Sections []struct {
		Start                   int     `json:"start"`
		Duration                float64 `json:"duration"`
		Confidence              int     `json:"confidence"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence int     `json:"time_signature_confidence"`
	} `json:"sections"`
	Segments []struct {
		Start           float64   `json:"start"`
		Duration        float64   `json:"duration"`
		Confidence      float64   `json:"confidence"`
		LoudnessStart   float64   `json:"loudness_start"`
		LoudnessMax     float64   `json:"loudness_max"`
		LoudnessMaxTime float64   `json:"loudness_max_time"`
		LoudnessEnd     int       `json:"loudness_end"`
		Pitches         []float64 `json:"pitches"`
		Timbre          []float64 `json:"timbre"`
	} `json:"segments"`
	Tatums []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"tatums"`
}
