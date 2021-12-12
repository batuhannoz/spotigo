package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type ArtistsTopTracksResponse struct {
	Tracks []struct {
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
	} `json:"tracks"`
}

func ArtistsTopTracks(input spotify.UserInfo) (ArtistsTopTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response ArtistsTopTracksResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artists/" + input.Id + "/top-tracks"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
