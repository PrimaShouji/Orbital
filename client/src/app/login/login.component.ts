import { Component, OnInit } from "@angular/core";
import { OrbitalServerService } from "../orbital/orbital-server.service";
import { Router } from "@angular/router";

@Component({
	selector: "app-login",
	templateUrl: "./login.component.html",
	styleUrls: ["./login.component.scss"],
})
export class LoginComponent implements OnInit {
	constructor(private orbital: OrbitalServerService, private router: Router) {}

	ngOnInit(): void {
		this.orbital.isAuthenticated().subscribe((res) => {
			if (res.is_authenticated) {
				this.router.navigate(["home"]);
			}
		});
	}
}
