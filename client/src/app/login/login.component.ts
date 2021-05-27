import { Component, OnInit } from "@angular/core";
import { OrbitalServerService } from "../orbital/orbital-server.service";

@Component({
	selector: "app-login",
	templateUrl: "./login.component.html",
	styleUrls: ["./login.component.scss"],
})
export class LoginComponent implements OnInit {
	authenticationStatus = "Not authenticated";

	constructor(private orbital: OrbitalServerService) {}

	ngOnInit(): void {
		this.orbital.isAuthenticated().subscribe(() => {
			this.authenticationStatus = "Authenticated";
		});
	}
}
