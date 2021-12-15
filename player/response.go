package spotify

type GetAvailableDevicesResponse struct {
	Devices []struct {
		Id               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"devices"`
}

type GetPlaybackStateResponse struct {
	Device struct {
		Id               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"device"`
	RepeatState  string `json:"repeat_state"`
	ShuffleState string `json:"shuffle_state"`
	Context      struct {
		Type         string `json:"type"`
		Href         string `json:"href"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Uri string `json:"uri"`
	} `json:"context"`
	Timestamp  int  `json:"timestamp"`
	ProgressMs int  `json:"progress_ms"`
	IsPlaying  bool `json:"is_playing"`
	Item       struct {
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
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		InterruptingPlayback  bool `json:"interrupting_playback"`
		Pausing               bool `json:"pausing"`
		Resuming              bool `json:"resuming"`
		Seeking               bool `json:"seeking"`
		SkippingNext          bool `json:"skipping_next"`
		SkippingPrev          bool `json:"skipping_prev"`
		TogglingRepeatContext bool `json:"toggling_repeat_context"`
		TogglingShuffle       bool `json:"toggling_shuffle"`
		TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
		TransferringPlayback  bool `json:"transferring_playback"`
	} `json:"actions"`
}

type GetRecentlyPlayedTracksResponse struct {
	Href  string `json:"href"`
	Items []struct {
	} `json:"items"`
	Limit   int    `json:"limit"`
	Next    string `json:"next"`
	Cursors struct {
		After string `json:"after"`
	} `json:"cursors"`
	Total int `json:"total"`
}

type GetCurrentlyPlayingResponse struct {
	Timestamp int64 `json:"timestamp"`
	Context   struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"context"`
	ProgressMs int `json:"progress_ms"`
	Item       struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			Id     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			Uri                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"artists"`
		DiscNumber  int  `json:"disc_number"`
		DurationMs  int  `json:"duration_ms"`
		Explicit    bool `json:"explicit"`
		ExternalIds struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		Id          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		IsPlayable  bool   `json:"is_playable"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewUrl  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		Uri         string `json:"uri"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Resuming bool `json:"resuming"`
		} `json:"disallows"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}
